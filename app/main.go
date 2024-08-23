package main

import (
	"BMI_calculator/service"
	"fmt"
)

func main() {
	bmiService := service.NewBMIService()

	var height, weight float64
	fmt.Print("Enter height in centimeters: ")
	fmt.Scanln(&height)
	fmt.Print("Enter weight in kilograms: ")
	fmt.Scanln(&weight)

	bmi := bmiService.CalculateBMI(height, weight)
	fmt.Printf("Your BMI is: %.2f\n", bmi)
}
