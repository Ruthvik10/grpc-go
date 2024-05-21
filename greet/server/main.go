package main

import (
	"log"
	"net"

	pb "github.com/Ruthvik10/grpc-go/greet/proto"
	"google.golang.org/grpc"
)

var (
	addr = "0.0.0.0:50051"
)

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v\n", addr, err)
	}
	defer lis.Close()

	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &handler{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
