package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:50007"
)

// 原生grpc client 负载均衡调用
// 通过dapr grpc proxy特性
// https://docs.dapr.io/operations/configuration/preview-features
// https://docs.dapr.io/developing-applications/building-blocks/service-invocation/howto-invoke-services-grpc
func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	for i := 0; ; i++ {
		ctx := metadata.AppendToOutgoingContext(context.Background(), "dapr-app-id", fmt.Sprintf("server%d", i%2))
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: fmt.Sprintf("Darth Tyrannus=>%d", i)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		log.Printf("Greeting: %s", r.GetMessage())
		time.Sleep(time.Second)
	}
}
