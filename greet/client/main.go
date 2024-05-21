package main

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	pb "github.com/Ruthvik10/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	// doGreet(client)
	// doGreetMany(client)
	// doLongGreet(client)
	doGreetEveryone(client)

}
func doGreet(client pb.GreetServiceClient) {
	res, err := client.Greet(context.Background(), &pb.GreetRequest{FirstName: "Micheal"})
	if err != nil {
		log.Println("Error recieving response:", err)
		return
	}
	log.Printf("Recieved response: %s", res.Result)
}

func doGreetMany(client pb.GreetServiceClient) {
	stream, err := client.GreetMany(context.Background(), &pb.GreetRequest{FirstName: "Micheal"})
	if err != nil {
		log.Println("Error recieving response:", err)
		return
	}
	for {
		msg, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Fatalln("Error while reading stream:", err)
		}

		log.Println(msg.Result)
	}
}

func doLongGreet(client pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest{
		{FirstName: "Micheal"},
		{FirstName: "Jim"},
		{FirstName: "Dwight"},
		{FirstName: "Andy"},
	}
	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalln("Error while calling LongGreet:", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error while recieving response", err)
	}

	log.Println(res.Result)
}

func doGreetEveryone(client pb.GreetServiceClient) {
	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalln("Error calling GreetEveryone:", err)
	}
	reqs := []*pb.GreetRequest{
		{FirstName: "Micheal"},
		{FirstName: "Jim"},
		{FirstName: "Dwight"},
		{FirstName: "Andy"},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(req.Result)
		}
		close(waitc)
	}()
	<-waitc
}
