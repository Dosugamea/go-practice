package main

import (
	"fmt"

	bmi "github.com/Dosugamea/go-practice/bmi/bmi_calculator"
)

func main() {
	if result, err := bmi.GetBmiFromHeightAndWeight(60.0, 47.0); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("BMI: %v\n", result)
	}
}
