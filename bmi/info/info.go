package info

import "fmt"


var mainTitle  = "BMI Calculator"
var separator  = "------------------------------------------"
var WeightText = "Please enter your weight (Kg): "
var HeightText = "Please enter your height (m): "

func PrintWelcome () {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}