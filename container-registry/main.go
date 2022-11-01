package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	crpb "github.com/x893675/malenia/proto/cr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

var (
	port = flag.Int("port", 5000, "The server port")
)

const (
	defaultStoreName = "statestore"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	crpb.UnimplementedHubServer
	client    dapr.Client
	storeName string
}

func NewServer() *server {
	client, err := dapr.NewClient()
	if err != nil {
		log.Panicln("FATAL! Dapr process/sidecar NOT found. Terminating!")
	}
	return &server{
		client:    client,
		storeName: defaultStoreName,
	}
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{grpc.MaxConcurrentStreams(10)}
	opts = append(opts)

	s := grpc.NewServer(opts...)

	srv := NewServer()
	crpb.RegisterHubServer(s, srv)

	log.Printf("Starting gRPC Server at %s", port)
	s.Serve(lis)
}

// Notice:
// grpc status code Ref: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
func (s server) CreateRepo(ctx context.Context, request *crpb.CreateRepoRequest) (*crpb.Repo, error) {
	if err := request.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	data, err := s.client.GetState(ctx, s.storeName, request.GetRepo().GetName(), nil)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if data.Value != nil {
		return nil, status.Error(codes.InvalidArgument, "repo has existed")
	}

	jsonPayload, _ := json.Marshal(request.GetRepo())

	request.Repo.Reset()
	if err := s.client.SaveState(ctx, s.storeName, request.GetRepo().GetName(), jsonPayload, nil); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return request.Repo, nil
}

func (s server) ListRepos(ctx context.Context, empty *emptypb.Empty) (*crpb.ListReposResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not impl")
}

func (s server) GetRepo(ctx context.Context, request *crpb.GetRepoRequest) (*crpb.Repo, error) {
	if err := request.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	data, err := s.client.GetState(ctx, s.storeName, request.GetName(), nil)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := &crpb.Repo{}
	if data.Value != nil {
		_ = json.Unmarshal(data.Value, resp)
	}
	return resp, nil
}
