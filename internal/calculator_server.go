// Package internal provides the implementation of the Calculator service.
package internal

import (
	"Calculations/api"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CalculatorServer is a struct that implements the Calculator service.
// It provides methods for adding and dividing numbers.
type CalculatorServer struct {
	calculator.UnimplementedCalculatorServer
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
	if request.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Cannot divide by zero")
	}

	var result float64
	done := make(chan bool)

	go func() {
		result = request.GetA() / request.GetB()
		done <- true
	}()

	<-done
	return &calculator.DivideResponse{Result: result}, nil
}
