package main

import (
	"Calculations/api"
	"Calculations/internal"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error("Failed to sync logger", zap.Error(err))
		}
	}(logger)

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}
	serverRegistrar := grpc.NewServer()
	reflection.Register(serverRegistrar)

	calculatorServer := internal.NewCalculatorServer(logger)
	calculator.RegisterCalculatorServer(serverRegistrar, calculatorServer)
	if err := serverRegistrar.Serve(lis); err != nil {
		logger.Fatal("Failed to serve", zap.Error(err))
	}
}
