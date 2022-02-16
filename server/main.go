package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

var (
	port = fmt.Sprintf(":%s", os.Args[1])
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 获取客户端metadata
	// map[:authority:[127.0.0.1:50051] content-type:[application/grpc+proto] dapr-app-id:[server0] grpc-trace-bin:[��      �ВR�}п��@!Zj��tp] user-agent:[grpc-go/1.40.0]]
	var md, ok = metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md)
	}

	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	for _, v := range os.Environ() {
		fmt.Println("env:", v)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
