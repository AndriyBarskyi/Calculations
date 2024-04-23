package main

import (
	"Calculations/api"
	"Calculations/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	serverRegistrar := grpc.NewServer()
	reflection.Register(serverRegistrar)

	calculator.RegisterCalculatorServer(serverRegistrar, &internal.CalculatorServer{})
	if err := serverRegistrar.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
