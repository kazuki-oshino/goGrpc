package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	pb "goGrpc/cmd/go_grpc"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The Server Port.")
)

type server struct {
	pb.UnimplementedExGoGrpcServer
}

func (s *server) HelloGrpc(ctx context.Context, in *pb.HelloGrpcRequest) (*pb.HelloGrpcResponse, error) {
	log.Printf("[name]Received Hello Grpc: %v\n", in.GetName())
	sgreets := make([]string, len(in.Greets))
	for i, v := range in.Greets {
		log.Printf("[greets]Received Hello Grpc: %v\n", v)
		sgreets[i] = v.String()
	}
	greets := strings.Join(sgreets, ",")

	return &pb.HelloGrpcResponse{Reply: fmt.Sprintf("Hello %v!!!", in.GetName()), Reply2: greets}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExGoGrpcServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
