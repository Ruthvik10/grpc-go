package main

import (
	"log"

	pb "github.com/Ruthvik10/grpc-go/calculator/proto"
)

type primeHandler struct {
	pb.UnimplementedPrimeServiceServer
}

func (h *primeHandler) Prime(in *pb.PrimeRequest, stream pb.PrimeService_PrimeServer) error {
	log.Println("Invoked Prime with argument:", in)
	var k int32 = 2
	n := in.Number
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeReponse{Result: k})
			n = n / k
		} else {
			k = k + 1
		}

	}
	return nil
}
