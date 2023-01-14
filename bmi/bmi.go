package main

import (
	"bmi/info"
)

func main() {
	// header
	info.PrintWelcome()
	
	// ask user for the information
	weight, height := GetUserMetrics()

	// calculate BMI
	bmi := calculateBMI(weight, height)
		
	printBMI(bmi)
}

func calculateBMI (weight float64, height float64) float64 {
	return weight / (height * height)
}
