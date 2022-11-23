package main

import "fmt"

func printBMI(bmi float64)  {
	if bmi <= 18.5 {
		fmt.Printf("Your BMI: %.2f \n", bmi)
		fmt.Printf("Underweight \n")
	}
	if bmi >= 18.5 && bmi < 25 {
		fmt.Printf("Your BMI: %.2f \n", bmi)
		fmt.Printf("Healthy Weight \n")
	}
	if bmi >= 25 && bmi < 30  {
		fmt.Printf("Your BMI: %.2f \n", bmi)
		fmt.Printf("Overweight \n")
	}
	if bmi >= 30 {
		fmt.Printf("Your BMI: %.2f \n", bmi)
		fmt.Printf("Obesity \n")
	}
}