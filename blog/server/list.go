package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("ListBlogs was invoked")
	cursor, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		data := &BlogItem{}
		err := cursor.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}
		stream.SendMsg(documentToBlog(data))
	}
	if err = cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil

}
