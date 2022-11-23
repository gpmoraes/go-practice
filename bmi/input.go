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
<<<<<<< HEAD
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
=======
var Reader = bufio.NewReader(os.Stdin)
>>>>>>> 060ceb58e133972e9052831e6f3ebe30ec7cd3ff
