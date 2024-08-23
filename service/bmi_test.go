package service

import (
	"fmt"
	"testing"
)

func TestCalculateBMI(t *testing.T) {
	// Create a new BMI service
	bmiService := NewBMIService()

	// Define test cases
	tests := []struct {
		height      float64
		weight      float64
		expectedBMI string
	}{
		{172, 69, "23.32"},
		{180, 80, "24.69"},
		{160, 50, "19.53"},
		{0, 50, "0.00"},    // Edge case: Height is zero, expect BMI to be zero
		{165, 0, "0.00"},   // Edge case: Weight is zero, expect BMI to be zero
		{165, 55, "20.20"}, // Normal case
	}

	for _, tt := range tests {
		// Calculate BMI
		bmi := bmiService.CalculateBMI(tt.height, tt.weight)

		// Format the BMI result to a string with 2 decimal places for comparison
		bmiStr := fmt.Sprintf("%.2f", bmi)

		// Check if the result matches the expected BMI (formatted to 2 decimal places)
		if bmiStr != tt.expectedBMI {
			t.Errorf("CalculateBMI(%v, %v) = %v; expected %v", tt.height, tt.weight, bmiStr, tt.expectedBMI)
		}
	}
}
