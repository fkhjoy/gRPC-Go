package main

import (
	"context"
	"log"

	pb "github.com/fkhjoy/gRPC-Go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	})

	if err != nil {
		log.Fatalf("Couldn't sum, %v", err)
	}

	log.Printf("Sum: %v\n", res.Result)
}
