package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/Ruthvik10/grpc-go/greet/proto"
)

type handler struct {
	pb.UnimplementedGreetServiceServer
}

func (h *handler) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Recieved message: %+v", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (h *handler) GreetMany(in *pb.GreetRequest, stream pb.GreetService_GreetManyServer) error {
	log.Printf("GreetMany was invoked with %v\n", in)
	for i := 1; i < 10; i++ {
		stream.Send(&pb.GreetResponse{Result: fmt.Sprintf("Hello %s, number %d", in.FirstName, i)})
	}
	return nil
}

func (h *handler) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet was invoked")
	res := ""
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.GreetResponse{Result: res})
			break
		}
		if err != nil {
			log.Fatalln("Error reading client stream:", err)
		}
		res += fmt.Sprintf("Hello %v\n", msg.FirstName)
	}
	return nil
}

func (h *handler) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked!")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln("Error recieving the client request:", err)
		}
		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalln("Error sending the response:", err)
		}
	}
}
