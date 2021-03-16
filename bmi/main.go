package main

import (
	"fmt"

	bmi "github.com/Dosugamea/go-practice/bmi/bmi_calculator"
)

func main() {
	var height bmi.Height = 160.0
	var weight bmi.Weight = 47.0
	if result, err := bmi.GetBmiFromHeightAndWeight(height, weight); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("BMI: %v\n", result)
	}
}
