package main

import (
	"context"
	"io"
	"log"

	pb "github.com/fkhjoy/gRPC-Go/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("do Greet many times")

	req := &pb.GreetRequest{
		FirstName: "Foysal",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}