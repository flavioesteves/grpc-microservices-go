package main

import (
	"context"
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("CalculatoSum was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
