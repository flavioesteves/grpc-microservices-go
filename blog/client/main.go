package main

import (
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)            // valid
	readBlog(c, "anInvalidID") // Invalid
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}

