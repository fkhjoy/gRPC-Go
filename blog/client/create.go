package main

import (
	"context"
	"log"

	pb "github.com/fkhjoy/gRPC-Go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("----createBlog was invoked----")

	blog := &pb.Blog{
		AuthorId: "Foysal",
		Title: "My first blog",
		Content: "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalln("Unexpected error: %v", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}