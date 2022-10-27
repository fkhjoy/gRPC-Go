package main

import (
	"context"
	"log"
	"time"

	pb "github.com/fkhjoy/gRPC-Go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest {
		{ FirstName: "Foysal"},
		{ FirstName: "Khandakar"},
		{ FirstName: "Joy"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalln("Error while calling longGreet: %v", err)
	}

	for _, req := range reqs {
		log.Println("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(1*time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)

	}

	log.Println("LongGreet: %s", res.Result)
}