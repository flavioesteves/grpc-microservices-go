package main

import (
	"context"
	"io"
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("primes was invoked")

	stream, err := c.Primes(context.Background(), &pb.PrimeRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("Primes: %d", msg.Result)
	}

}
