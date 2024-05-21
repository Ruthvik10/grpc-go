package main

import (
	"context"
	"log"
	"math"

	pb "github.com/Ruthvik10/grpc-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type sqrtHandler struct {
	*pb.UnimplementedSqrtServiceServer
}

func (h *sqrtHandler) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Println("Recieved number:", in.Num)
	if in.Num < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Recieved a negative number: %d", in.Num)
	}
	sqrt := math.Sqrt(float64(in.Num))
	return &pb.SqrtResponse{Res: sqrt}, nil
}
