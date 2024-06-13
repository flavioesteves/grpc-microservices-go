package main

import (
	"context"
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/calculator/proto"
)

func (s *Server) CalculatorSum(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("CalculatoSum was invoked with %v\n", in)
	return &pb.CalculatorResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
