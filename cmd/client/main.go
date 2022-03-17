package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "goGrpc/cmd/go_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewExGoGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	greets := make([]pb.Greet, 3)
	greets[0] = pb.Greet_HELLO
	greets[1] = pb.Greet_WHATS_UP
	greets[2] = pb.Greet_HELLO

	r, err := c.HelloGrpc(ctx, &pb.HelloGrpcRequest{Name: "Kazuki", Greets: greets})

	if err != nil {
		log.Fatalf("could not hello: %v", err)
	}
	log.Printf("Greeting: %s\n", r.GetReply())
	log.Printf("Greeting: %s", r.GetReply2())
}
