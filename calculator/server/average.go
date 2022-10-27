package main

import (
	"io"
	"log"

	pb "github.com/fkhjoy/gRPC-Go/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average was invoked")

	var sum int64 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			avg := float64(sum)/float64(count)
			return stream.SendAndClose(&pb.AvgResponse{
				Result: avg,
			})
		}

		if err != nil {
			log.Fatalln("Error while reading client stream: %v", err)
		}
		
		log.Printf("Receiving req: %v\n", req)
		sum += req.Number;
		count++;
	}

}