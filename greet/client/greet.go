package main

import (
	"context"
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoker")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Flavio",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}
