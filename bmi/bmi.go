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
		
<<<<<<< HEAD
	printBMI(bmi)
}

func calculateBMI (weight float64, height float64) float64 {
	return weight / (height * height)
}
=======
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
>>>>>>> 060ceb58e133972e9052831e6f3ebe30ec7cd3ff
