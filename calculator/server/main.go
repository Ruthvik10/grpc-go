package main

import (
	"log"
	"net"

	pb "github.com/Ruthvik10/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

var addr = "localhost:50052"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}
	defer lis.Close()

	log.Println("Listening on", addr)

	s := grpc.NewServer()

	pb.RegisterSumServiceServer(s, &sumHandler{})
	pb.RegisterPrimeServiceServer(s, &primeHandler{})
	pb.RegisterAvgServiceServer(s, &avgHandler{})
	pb.RegisterSqrtServiceServer(s, &sqrtHandler{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to server: ", err)
	}
}
