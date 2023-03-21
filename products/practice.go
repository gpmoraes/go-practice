// Package main is the entry point for this application.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Product represents a product with its details such as ID, title, description and price.
type Product struct {
	id          string
	title       string
	description string
	price       float64
}

// Store stores the product details in a text file with the filename as the product ID.
func (prod *Product) store() {
	file, _ := os.Create(prod.id + ".txt")

	content := fmt.Sprintf(
		"ID: %v\nTITLE: %v\nDSC: %v\nPrice: USD %.2f\n",
		prod.id,
		prod.title,
		prod.description,
		prod.price,
	)

	file.WriteString(content)

	file.Close()
}

// PrintData prints the product details on the console.
func (prod *Product) printData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("TITLE: %v\n", prod.title)
	fmt.Printf("DSC: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n", prod.price)
}

// NewProduct returns a new product object with the given parameters.
func NewProduct(id string, t string, d string, p float64) *Product {
	return &Product{id, t, d, p}
}

// Main function of the application.
func main() {
	// Get the product details from the user and create a new product object.
	createdProduct := getProduct()

	// Print the product details on the console.
	createdProduct.printData()

	// Store the product details in a text file.
	createdProduct.store()
}

// getProduct gets the product details from the user and returns a new product object.
func getProduct() *Product {
	fmt.Println("Please enter the product data.")
	fmt.Println("-------------------------------------")

	reader := bufio.NewReader(os.Stdin)

	// Read the product details from the user.
	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product Title: ")
	dscInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")

	// Convert the price string to a float value.
	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	// Create a new product object with the read details.
	product := NewProduct(
		idInput,
		titleInput,
		dscInput,
		priceValue,
	)

	return product
}

// readUserInput reads the user input from the console and returns the value.
func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)

	// Read the user input from the console.
	userInput, _ := reader.ReadString('\n')

	// Remove the newline character from the input string.
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
