package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	daprhttp "github.com/dapr/go-sdk/service/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	port = flag.Int("port", 5000, "The server port")
)

func main() {
	log.Printf("### Server listening on %v\n", *port)

	router := mux.NewRouter()
	AddHealth(router)

	s := NewService(router)

	router.HandleFunc("/core.kubeclipper.io/v1/clusters", s.CreateCluster).Methods("POST")
	router.HandleFunc("/core.kubeclipper.io/v1/clusters/{name}", s.DeleteCluster).Methods("DELETE")

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", *port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func AddHealth(r *mux.Router) {
	// Add health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})
}

const (
	defaultStoreName = "statestore"
	//defaultPubSub    = "pubsub"
	//defaultTopic     = "clusters-queue"
)

type Service struct {
	storeName string
	client    dapr.Client
}

func NewService(router *mux.Router) *Service {
	client, err := dapr.NewClient()
	if err != nil {
		log.Panicln("FATAL! Dapr process/sidecar NOT found. Terminating!")
	}
	s := &Service{
		storeName: defaultStoreName,
		client:    client,
	}
	// We don't actually use the service as we have one already
	// But we need to call AddTopicEventHandler to register the handler
	dummyService := daprhttp.NewServiceWithMux("notUsed", router)

	_ = dummyService.Start()

	return s
}

type ClusterStatus string

const (
	ClusterInstalling  ClusterStatus = "Installing"
	ClusterRunning     ClusterStatus = "Running"
	ClusterTerminating ClusterStatus = "Terminating"
)

type Cluster struct {
	Name   string        `json:"name"`
	Status ClusterStatus `json:"status,omitempty"`
}

func (s *Service) CreateCluster(resp http.ResponseWriter, req *http.Request) {
	c := Cluster{}
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.SetStatus(req.Context(), &c, ClusterInstalling)
	if err != nil {
		log.Printf("### Failed to update state for cluster %s\n", err)
		resp.WriteHeader(http.StatusInternalServerError)
		_, _ = resp.Write([]byte(err.Error()))
		return
	}

	go func() {
		// Fake background cluster completion
		time.AfterFunc(30*time.Second, func() {
			log.Printf("### Cluster %s is now Installing Successful\n", c.Name)
			_ = s.SetStatus(context.Background(), &c, ClusterRunning)
		})
	}()

	cBytes, _ := json.Marshal(c)
	_, _ = resp.Write(cBytes)
}

func (s *Service) SetStatus(ctx context.Context, c *Cluster, status ClusterStatus) error {
	log.Printf("### Setting status for cluster %s to %s\n", c.Name, status)
	c.Status = status

	// Save updated order list back, again keyed using user id
	jsonPayload, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("err://json-marshall, State JSON marshalling error, service: kc-server, %s", err.Error())
	}

	if err := s.client.SaveState(ctx, s.storeName, c.Name, jsonPayload, nil); err != nil {
		log.Printf("### Error! Unable to update status of cluster '%s'", c.Name)
		return fmt.Errorf("dapr status problem , service: kc-server,  %s", err.Error())
	}

	return nil
}

func (s *Service) DeleteCluster(resp http.ResponseWriter, req *http.Request) {
	parms := mux.Vars(req)
	cluName := parms["name"]
	data, err := s.client.GetState(req.Context(), s.storeName, cluName, nil)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		_, _ = resp.Write([]byte(err.Error()))
		return
	}
	if data.Value == nil {
		resp.WriteHeader(http.StatusBadRequest)
		_, _ = resp.Write([]byte("not found"))
		return
	}
	c := Cluster{}
	_ = json.Unmarshal(data.Value, &c)
	if c.Status == ClusterInstalling {
		resp.WriteHeader(http.StatusBadRequest)
		_, _ = resp.Write([]byte("can not delete when cluster is installing"))
		return
	}
	if err := s.client.DeleteState(req.Context(), s.storeName, c.Name, nil); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		_, _ = resp.Write([]byte(err.Error()))
		return
	}
	resp.WriteHeader(http.StatusOK)
}
