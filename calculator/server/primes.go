package main

import (
	"log"

	pb "github.com/flavioesteves/grpc-microservices-go/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {

	log.Printf("Primes function was invoked with %v\n", in)

	k := int64(2)
	n := in.Number

	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			n /= k
		} else {
			k++
		}
	}

	return nil
}
