package main

import (
	"context"
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("-----deleteBlog was invoked-----")

	blog := pb.BlogId{
		Id: id,
	}

	_, err := c.DeleteBlog(context.Background(), &blog)

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Printf("Blog was deleted %s\n", id)

}
