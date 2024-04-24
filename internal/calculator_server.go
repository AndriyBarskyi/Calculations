// Package internal provides the implementation of the Calculator service.
package internal

import (
	"Calculations/api"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CalculatorServer is a struct that implements the Calculator service.
// It provides methods for adding and dividing numbers.
type CalculatorServer struct {
	calculator.UnimplementedCalculatorServer
	logger *zap.Logger
}

func NewCalculatorServer(logger *zap.Logger) *CalculatorServer {
	return &CalculatorServer{logger: logger}
}

// Add is a method of CalculatorServer that takes a context and an AddRequest as parameters.
// It returns an AddResponse and an error.
// The method adds the two numbers provided in the request and returns the result in the response.
// If an error occurs during the operation, it returns the error.
func (c *CalculatorServer) Add(_ context.Context, request *calculator.AddRequest) (*calculator.AddResponse, error) {
	var result float64
	done := make(chan bool)

	go func() {
		result = request.GetA() + request.GetB()
		done <- true
	}()

	<-done
	return &calculator.AddResponse{Result: result}, nil
}

// Divide is a method of CalculatorServer that takes a context and a DivideRequest as parameters.
// It returns a DivideResponse and an error.
// The method divides the first number by the second number provided in the request and returns the result in the response.
// If the second number is zero, it returns an error indicating that division by zero is not allowed.
// If an error occurs during the operation, it returns the error.
func (c *CalculatorServer) Divide(_ context.Context, request *calculator.DivideRequest) (*calculator.DivideResponse, error) {
	c.logger.Info("Divide called", zap.Float64("A", request.GetA()), zap.Float64("B", request.GetB()))

	if request.GetB() == 0 {
		err := status.Error(codes.InvalidArgument, "Cannot divide by zero")
		c.logger.Error("Error in Divide", zap.Error(err))
		return nil, err
	}

	var result float64
	done := make(chan bool)

	go func() {
		result = request.GetA() / request.GetB()
		done <- true
	}()

	<-done
	c.logger.Info("Divide result", zap.Float64("Result", result))
	return &calculator.DivideResponse{Result: result}, nil
}
