package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Ruthvik10/grpc-go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var addr = "localhost:50052"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed to connect:", err)
	}
	defer conn.Close()

	// client := pb.NewSumServiceClient(conn)
	// doSum(client)

	// client2 := pb.NewPrimeServiceClient(conn)
	// doPrime(client2)

	// client3 := pb.NewAvgServiceClient(conn)
	// doAverage(client3)

	client4 := pb.NewSqrtServiceClient(conn)
	doSqrt(client4, -10)
}

func doSum(client pb.SumServiceClient) {
	res, err := client.Sum(context.Background(), &pb.SumRequest{FirstNum: 10, SecondNum: 30})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Sum =", res.Sum)
}

func doPrime(client pb.PrimeServiceClient) {
	stream, err := client.Prime(context.Background(), &pb.PrimeRequest{Number: 120})
	if err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error recieving values:", err)
		}

		log.Println("Recieved: ", msg.Result)
	}
}

func doAverage(client pb.AvgServiceClient) {
	stream, err := client.Average(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	reqs := []*pb.AvgRequest{{Num: 1}, {Num: 2}, {Num: 10}}

	for _, req := range reqs {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error recieving response from the from the server", err)
	}
	log.Println(res.Avg)
}

func doSqrt(client pb.SqrtServiceClient, n int32) {
	res, err := client.Sqrt(context.Background(), &pb.SqrtRequest{Num: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Println(e.Message())
			log.Println(e.Code())
			return
		}
		log.Fatal("Recieved non gRPC error")
	}
	log.Printf("Sqrt of %d is %f\n", n, res.Res)
}
