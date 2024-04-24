package internal

import (
	"Calculations/api"
	"context"
	"go.uber.org/zap"
	"testing"
)

func TestCalculatorServer_Add(t *testing.T) {
	server := &CalculatorServer{}

	testCases := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"zeroes", 0, 0, 0},
		{"negative numbers", -2, -3, -5},
		{"positive and negative numbers", 2, -3, -1},
		{"negative and positive numbers", -2, 3, 1},
		{"zero and positive number", 0, 3, 3},
		{"positive number and zero", 3, 0, 3},
		{"zero and negative number", 0, -3, -3},
		{"negative number and zero", -3, 0, -3},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // marks this test to be run in parallel
			req := &calculator.AddRequest{A: tc.a, B: tc.b}
			resp, err := server.Add(context.Background(), req)
			if err != nil {
				t.Fatalf("Add returned an error: %v", err)
			}
			if resp.Result != tc.expected {
				t.Errorf("Add(%v, %v) returned incorrect result: got %v, want %v", tc.a, tc.b, resp.Result, tc.expected)
			}
		})
	}
}

func TestCalculatorServer_Divide(t *testing.T) {
	server := &CalculatorServer{}
	logger, err := zap.NewProduction()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	testCases := []struct {
		name        string
		a, b        float64
		expected    float64
		expectError bool
	}{
		{"positive numbers", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"zero divided by number", 0, 2, 0, false},
		{"number divided by one", 10, 1, 10, false},
		{"negative numbers", -10, -2, 5, false},
		{"negative divided by positive", -10, 2, -5, false},
		{"positive divided by negative", 10, -2, -5, false},
		{"zero divided by zero", 0, 0, 0, true},
		{"negative divided by zero", -10, 0, 0, true},
		{"zero divided by negative", 0, -2, 0, false},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // marks this test to be run in parallel
			req := &calculator.DivideRequest{A: tc.a, B: tc.b}
			resp, err := server.Divide(context.Background(), req)
			if tc.expectError {
				if err == nil {
					logger.Error("Divide did not return an error when expected", zap.Float64("A", tc.a), zap.Float64("B", tc.b))
					t.Errorf("Divide(%v, %v) did not return an error when expected", tc.a, tc.b)
				}
			} else {
				if err != nil {
					logger.Error("Error in Divide", zap.Error(err))
					t.Fatalf("Divide returned an error: %v", err)
				}
				if resp.Result != tc.expected {
					logger.Error("Incorrect result in Divide", zap.Float64("Got", resp.Result), zap.Float64("Expected", tc.expected))
					t.Errorf("Divide(%v, %v) returned incorrect result: got %v, want %v", tc.a, tc.b, resp.Result, tc.expected)
				}
			}
		})
	}
}
