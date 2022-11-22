package main

import (
	"bmi/info"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// header
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	
	// ask user for the information
	fmt.Print(info.WeightText)
	// as the reader will return two values we can ignore the second one
	// adding a underline at the second variable
	inputWeight, _ := Reader.ReadString('\n')

	fmt.Print(info.HeightText)
	inputHeight, _ := Reader.ReadString('\n')

	// clear the input line break
	inputWeight = strings.Replace(inputWeight, "\n", "", -1)
	inputHeight = strings.Replace(inputHeight, "\n", "", -1)

	// turn the input (string) into float
	weight, _ := strconv.ParseFloat(inputWeight, 64)
	height, _ := strconv.ParseFloat(inputHeight, 64)

	// calculate BMI
	bmi := weight / (height * height)
		
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
