package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/fkhjoy/gRPC-Go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("do Greet Everyone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalln("Error while creating stream: %v", err)

	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Foysal"},
		{FirstName: "Khandakar"},
		{FirstName: "Joy"},
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