package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/fkhjoy/gRPC-Go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("do Max was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalln("Error while creating stream: %v", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Println("Send request: %v", req)

			stream.Send(req)
			time.Sleep(1*time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalln("Error while receving: %v", err)
				break
			}

			log.Println("Received: %v", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}