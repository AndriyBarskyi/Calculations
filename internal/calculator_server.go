// Package internal provides the implementation of the Calculator service.
package internal

import (
	calculator "Calculations/api"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CalculatorServer is a struct that implements the Calculator service.
// It provides methods for multiplying and dividing numbers.
type CalculatorServer struct {
	calculator.UnimplementedCalculatorServer
	logger *zap.Logger
}

func NewCalculatorServer(logger *zap.Logger) *CalculatorServer {
	return &CalculatorServer{logger: logger}
}

// Calculate is a method of CalculatorServer that takes a context and a CalculateRequest as parameters.
// It returns a CalculateResponse and an error.
// The method multiplies and divides the two numbers provided in the request and returns the results in the response.
// If the second number is zero, it returns an error indicating that division by zero is not allowed.
// If an error occurs during the operation, it returns the error.
func (c *CalculatorServer) Calculate(_ context.Context, request *calculator.CalculateRequest) (*calculator.CalculateResponse, error) {
	c.logger.Info("Calculate called", zap.Float64("A", request.GetA()), zap.Float64("B", request.GetB()))

	if request.GetB() == 0 {
		err := status.Error(codes.InvalidArgument, "Cannot divide by zero")
		c.logger.Error("Error in Calculate", zap.Error(err))
		return nil, err
	}

	multiplyResult := make(chan float64)
	divideResult := make(chan float64)

	go func() {
		multiplyResult <- request.GetA() * request.GetB()
	}()

	go func() {
		divideResult <- request.GetA() / request.GetB()
	}()

	return &calculator.CalculateResponse{MultiplyResult: <-multiplyResult, DivideResult: <-divideResult}, nil
}
