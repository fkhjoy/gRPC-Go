package main

import (
	"context"
	"log"
	"time"

	pb "github.com/fkhjoy/gRPC-Go/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage function was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 4},
		{Number: 5},
		{Number: 3},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalln("Error while calling Average: %v", err)
	}

	for _, req := range reqs {
		log.Println("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)

	}

	log.Println("Average: %v", res.Result)
}
