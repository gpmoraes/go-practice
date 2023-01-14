package main

import (
	"bmi/info"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// variable that stores the user input
var reader = bufio.NewReader(os.Stdin)

func GetUserMetrics() (weight float64, height float64) {
	weight = getUserInput(info.WeightText)
	height = getUserInput(info.HeightText)

	return
}

func getUserInput(promptText string) (value float64) {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	value, _  = strconv.ParseFloat(userInput, 64)

	return
}
