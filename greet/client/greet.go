package main

import (
	"context"
	"log"

	pb "github.com/fkhjoy/gRPC-Go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Foysal",
	})

	if err != nil {
		log.Fatalf("Couldn't greet, %v", err)
	}

	log.Printf("Greeting: %s\n", res.Result)

}
