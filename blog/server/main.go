package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fkhjoy/gRPC-Go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatalln(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Couldn't listen to %v", addr)
	}

	log.Printf("Listening to %v", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
