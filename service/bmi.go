package service

import (
	"github.com/shopspring/decimal"
)

type BMIService interface {
	CalculateBMI(height, weight float64) float64
}

type bmiService struct{}

func NewBMIService() BMIService {
	return &bmiService{}
}

// CalculateBMI BMI = weight / (height)^2
func (s *bmiService) CalculateBMI(height, weight float64) float64 {
	// Convert height from centimeters to meters
	heightInMeters := decimal.NewFromFloat(height).Div(decimal.NewFromFloat(100))
	weightInKg := decimal.NewFromFloat(weight)

	// Calculate BMI = weight / (height in meters)^2
	heightSquared := heightInMeters.Mul(heightInMeters)
	if heightSquared.IsZero() {
		return 0 // Prevent division by zero
	}

	bmi := weightInKg.Div(heightSquared)

	// Convert decimal result back to float64
	bmiFloat, _ := bmi.Float64()
	return bmiFloat
}
