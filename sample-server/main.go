package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	helloworldpb "github.com/x893675/malenia/proto/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 5000, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	helloworldpb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworldpb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	s := daprd.NewServiceWithListener(lis)

	_ = s.AddServiceInvocationHandler("echo", func(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
		log.Printf("in echo handler %v\n", *in)
		out = &common.Content{
			Data:        in.Data,
			ContentType: in.ContentType,
			DataTypeURL: in.DataTypeURL,
		}
		return
	})

	helloworldpb.RegisterGreeterServer(srv, &server{})

	pb.RegisterAppCallbackServer(srv, s.(*daprd.Server))

	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
