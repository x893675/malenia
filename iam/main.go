package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	envoytype "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"github.com/gogo/googleapis/google/rpc"
	iampb "github.com/x893675/malenia/proto/iam"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"strings"
)

var (
	port = flag.Int("port", 5000, "The server port")
)

var (
	defaultUserMap = map[string]*iampb.User{
		"spike": {
			Name:  "spike",
			Email: "spike@example.org",
		},
		"hanamichi": {
			Name:  "hanamichi",
			Email: "hanamichi@example.org",
		},
	}
)

type healthServer struct{}

func (s *healthServer) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	log.Printf("Handling grpc Check request, service %s", in.Service)
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		log.Println("no incoming context")
	} else {
		log.Printf("%v", md)
	}
	// yeah, right, open 24x7, like 7-11
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s *healthServer) Watch(in *healthpb.HealthCheckRequest, srv healthpb.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Watch is not implemented")
}

type AuthorizationServer struct {
	iampb.UnimplementedIdentityAccessManagementServer
}

func (a *AuthorizationServer) ListUsers(ctx context.Context, empty *emptypb.Empty) (*iampb.ListUsersResponse, error) {
	resp := &iampb.ListUsersResponse{}
	for _, v := range defaultUserMap {
		resp.Users = append(resp.Users, v)
	}
	return resp, nil
}

func (a *AuthorizationServer) Enforce(ctx context.Context, request *iampb.EnforceRequest) (*iampb.EnforceReply, error) {
	log.Printf("%#v\n", request)
	resp := &iampb.EnforceReply{
		Result: false,
		Reason: "",
	}
	if _, ok := defaultUserMap[request.Username]; ok {
		resp.Result = true
	} else {
		resp.Reason = fmt.Sprintf("user %s has no permission to visit url=%s action=%s method=%s",
			request.Username, request.Username, request.Action, request.Method)
	}
	return resp, nil
}

func (a *AuthorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	log.Println(">>> Authorization called check()")

	b, err := json.MarshalIndent(req.Attributes.Request.Http.Headers, "", "  ")
	if err == nil {
		log.Println("Inbound Headers: ")
		log.Println(string(b))
	}

	ct, err := json.MarshalIndent(req.Attributes.ContextExtensions, "", "  ")
	if err == nil {
		log.Println("Context Extensions: ")
		log.Println(string(ct))
	}

	authHeader, ok := req.Attributes.Request.Http.Headers["authorization"]
	var splitToken []string

	if ok {
		splitToken = strings.Split(authHeader, "Bearer ")
	}
	if len(splitToken) == 2 {
		token := splitToken[1]

		if _, ok := defaultUserMap[token]; ok {
			return &auth.CheckResponse{
				Status: &rpcstatus.Status{
					Code: int32(rpc.OK),
				},
				HttpResponse: &auth.CheckResponse_OkResponse{
					OkResponse: &auth.OkHttpResponse{
						Headers: []*core.HeaderValueOption{
							{
								Header: &core.HeaderValue{
									Key:   "x-custom-header-from-authz",
									Value: "some value",
								},
							},
						},
						HeadersToRemove: []string{
							"authorization",
						},
					},
				},
			}, nil
		} else {
			return &auth.CheckResponse{
				Status: &rpcstatus.Status{
					Code: int32(rpc.PERMISSION_DENIED),
				},
				HttpResponse: &auth.CheckResponse_DeniedResponse{
					DeniedResponse: &auth.DeniedHttpResponse{
						Status: &envoytype.HttpStatus{
							Code: envoytype.StatusCode_Unauthorized,
						},
						Body: "PERMISSION_DENIED",
					},
				},
			}, nil
		}
	}
	return &auth.CheckResponse{
		Status: &rpcstatus.Status{
			Code: int32(rpc.UNAUTHENTICATED),
		},
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Status: &envoytype.HttpStatus{
					Code: envoytype.StatusCode_Unauthorized,
				},
				Body: "Authorization Header malformed or not provided",
			},
		},
	}, nil
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

	authzserver := &AuthorizationServer{}
	auth.RegisterAuthorizationServer(s, authzserver)
	healthpb.RegisterHealthServer(s, &healthServer{})
	iampb.RegisterIdentityAccessManagementServer(s, authzserver)

	log.Printf("Starting gRPC Server at %s", port)
	s.Serve(lis)
}
