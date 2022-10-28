package main

import (
	"io"
	"log"

	pb "github.com/fkhjoy/gRPC-Go/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	var maxi int64
	var first bool = true
	change := true
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln("Error while reading client: %v", err)
		}

		n := req.Number
		if first {
			maxi = n
			first = false
			
		} else {
			if n >= maxi {
				maxi = n
				change = true
			}
		}
		if change {
			err = stream.Send(&pb.MaxResponse{
				Result: maxi,
			})
			if err != nil {
				log.Fatalln("Error while sending data: %v", err)
			}
			change = false
		}
		
	}
}
