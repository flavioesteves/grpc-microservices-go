package main

import (
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	connection, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer connection.Close()

	c := pb.NewCalculatorServiceClient(connection)
	//doSum(c)
	//doPrimes(c)
	//doAvg(c)
	//doMax(c)
	//doSqrt(c, 10)
	doSqrt(c, -2)
}
