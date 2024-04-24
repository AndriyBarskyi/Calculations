package internal

import (
	"Calculations/api"
	"context"
	"go.uber.org/zap"
	"testing"
)

func TestCalculatorServer_Calculate(t *testing.T) {
	logger, err := zap.NewProduction()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	server := NewCalculatorServer(logger)

	testCases := []struct {
		name        string
		a, b        float64
		expectedMul float64
		expectedDiv float64
		expectError bool
	}{
		{"positive numbers", 10, 2, 20, 5, false},
		{"division by zero", 10, 0, 0, 0, true},
		{"zero divided by number", 0, 2, 0, 0, false},
		{"number divided by one", 10, 1, 10, 10, false},
		{"negative numbers", -10, -2, 20, 5, false},
		{"negative divided by positive", -10, 2, -20, -5, false},
		{"positive divided by negative", 10, -2, -20, -5, false},
		{"zero divided by zero", 0, 0, 0, 0, true},
		{"negative divided by zero", -10, 0, 0, 0, true},
		{"zero divided by negative", 0, -2, 0, 0, false},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // marks this test to be run in parallel
			req := &calculator.CalculateRequest{A: tc.a, B: tc.b}
			resp, err := server.Calculate(context.Background(), req)
			if tc.expectError {
				if err == nil {
					logger.Error("Calculate did not return an error when expected", zap.Float64("A", tc.a), zap.Float64("B", tc.b))
					t.Errorf("Calculate(%v, %v) did not return an error when expected", tc.a, tc.b)
				}
			} else {
				if err != nil {
					logger.Error("Error in Calculate", zap.Error(err))
					t.Fatalf("Calculate returned an error: %v", err)
				}
				if resp.MultiplyResult != tc.expectedMul || resp.DivideResult != tc.expectedDiv {
					logger.Error("Incorrect result in Calculate", zap.Float64("Got Multiply", resp.MultiplyResult), zap.Float64("Expected Multiply", tc.expectedMul), zap.Float64("Got Divide", resp.DivideResult), zap.Float64("Expected Divide", tc.expectedDiv))
					t.Errorf("Calculate(%v, %v) returned incorrect result: got Multiply %v, want Multiply %v, got Divide %v, want Divide %v", tc.a, tc.b, resp.MultiplyResult, tc.expectedMul, resp.DivideResult, tc.expectedDiv)
				}
			}
		})
	}
}
