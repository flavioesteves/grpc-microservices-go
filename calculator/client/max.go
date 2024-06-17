package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/flavioesteves/grpc-microservices-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	//Go Routine to send
	go func() {

		numbers := []int32{1, 5, 3, 6, 3, 20}

		for _, number := range numbers {
			log.Printf("Send request: %d\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// Go Routune to receive from the server
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received a new maximum: %d\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
