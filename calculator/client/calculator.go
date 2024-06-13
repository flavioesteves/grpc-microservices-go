package main

import (
	"context"
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/calculator/proto"
)

func doSumCalculator(c pb.CalculatorSumServiceClient) {
	log.Println("doSumCalculator was invoked")

	res, err := c.CalculatorSum(context.Background(), &pb.CalculatorRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("Result: %v\n", res.Result)
}
