package main

import (
	"io"
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {

	log.Println("Avg function was invoked")

	var sum int64 = 0
	count := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: calculateAvg(int64(sum), int64(count)),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		sum += req.Number
		count++

		log.Printf("Receiving: %v\n", req.Number)
	}
}

func calculateAvg(sum int64, count int64) float64 {
	return float64(float64(sum) / float64(count))
}
