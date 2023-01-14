# BMI Calculator
This is a simple command-line application that calculates Body Mass Index (BMI) based on user input. The application prompts the user to enter their weight in kilograms and height in meters, and then calculates the BMI using the formula weight/(height*height). The application then classifies the BMI into one of four categories: underweight, healthy weight, overweight, and obesity.

## Concepts used in this project
- **Packages**: The application is organized into multiple packages, each with a specific responsibility. The `main` package contains the entry point of the application, the `info` package contains the texts and message to be shown to the user, the `input` package is responsible for receiving the user input and the `output` package is responsible for showing the results to the user
- **Variables**: The application uses variables to store user input, calculated results, and text prompts.
- **Functions**: The application uses functions to organize and structure the code. The main function is `main()`, which calls other functions to perform specific tasks such as receiving user input, calculating the BMI, and printing the results.
- **Control Flow**: The application uses control flow statements (`if` and `else`) to determine which category the BMI belongs to and then print the appropriate message to the user.

## How to use
1. Make sure you have Go installed on your machine
2. Clone the repository to your local machine
3. Navigate to the project folder using the command line
4. Run the command `go run bmi.go`
5. Follow the prompts to enter your weight and height
6. The application will calculate your BMI and print the results to the screen

## Tips
- Make sure you enter your weight in kilograms and your height in meters, otherwise the calculation will not be accurate.
- If you wish to change the texts that the application uses, you can do it in the info.go file, where you will find all the texts used by the application.
