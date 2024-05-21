package main

import (
	"io"
	"log"

	pb "github.com/Ruthvik10/grpc-go/calculator/proto"
)

type avgHandler struct {
	*pb.UnimplementedAvgServiceServer
}

func (h *avgHandler) Average(stream pb.AvgService_AverageServer) error {
	var avg float64
	var sum, count = 0.0, 0.0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			avg = sum / count
			return stream.SendAndClose(&pb.AvgResponse{Avg: avg})
		}
		if err != nil {
			log.Fatalln("Error while recieving the client stream:", err)
		}
		count++
		sum += float64(msg.Num)

	}
}
