package main

import (
	"context"
	"log"

	pb "github.com/Ruthvik10/grpc-go/calculator/proto"
)

type sumHandler struct {
	pb.UnimplementedSumServiceServer
}

func (h *sumHandler) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Recieved arguments:", in)
	return &pb.SumResponse{
		Sum: in.FirstNum + in.SecondNum,
	}, nil
}
