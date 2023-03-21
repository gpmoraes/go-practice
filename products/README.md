# Product Management System

This is a simple command-line application that allows the user to input data for a product, store it in a text file, and print the information to the console. The application uses a `Product` struct to represent the product data, and includes methods to print the data and store it in a file.

## Concepts used in this project
- **Structs**: The `Product` struct is used to represent the product data, and includes fields for the product ID, title, description, and price.
- **Methods**: The `Product` struct includes two methods, `printData()` and `store()`, which print the product data to the console and store it in a file, respectively.
- **File I/O**: The `store()` method uses file I/O to create a text file with the product data, which is saved in the same directory as the application.
- **User Input**: The `getProduct()` function prompts the user to input data for the product, and returns a `Product` struct with the user's input.

## How to use
1. Make sure you have Go installed on your machine.
2. Clone the repository to your local machine.
3. Navigate to the project folder using the command line.
4. Run the command `go run main.go`.
5. Follow the prompts to enter the product data.
6. The application will print the product data to the console and create a text file with the data.

## Tips
- Make sure to enter a unique ID for each product, as the ID is used to name the text file that stores the product data.
- If you wish to modify the fields in the `Product` struct, you can do so in the `NewProduct()` function.
