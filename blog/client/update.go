package main

import (
	"context"
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked ---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Mac",
		Title:    "A new title",
		Content:  "Content of the first blog, with some additions!",
	}

	_, err := c.Updateblog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updateing: %v\n", err)
	}
	log.Println("Blog was updated")
}
