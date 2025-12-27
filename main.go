package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to my calculator program.")
	fmt.Println("Which operation would you like to perform today?")
	fmt.Println("a. add")
	fmt.Println("b. subtract")
	fmt.Println("c. multiply")
	fmt.Println("d. divide")
	fmt.Println("e. absolute value")
	fmt.Println("f. round")
	fmt.Println("g. raise to an exponent")
	fmt.Println("h. square root")

	var choice string
	fmt.Print("What operation do you want to choose: ")
	fmt.Scanln(&choice)

	var num1, num2 float64

	switch choice {
	case "a":
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)

	case "b":
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)

	case "c":
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)

	case "d":
		fmt.Print("Enter dividend: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter divisor: ")
		fmt.Scanln(&num2)
		fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)

	case "e":
		fmt.Print("Enter a number: ")
		fmt.Scanln(&num1)
		fmt.Printf("The absolute value of %.2f is %.2f\n", num1, math.Abs(num1))

	case "f":
		fmt.Print("Enter a number: ")
		fmt.Scanln(&num1)
		fmt.Printf("%.2f rounded is %.0f\n", num1, math.Round(num1))

	case "g":
		fmt.Print("Enter base number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter exponent: ")
		fmt.Scanln(&num2)
		fmt.Printf("%.2f raised to the power of %.2f is %.2f\n", num1, num2, math.Pow(num1, num2))

	case "h":
		fmt.Print("Enter a non-negative number: ")
		fmt.Scanln(&num1)
		fmt.Printf("The square root of %.2f is %.2f\n", num1, math.Sqrt(num1))

	default:
		fmt.Println("Invalid choice.")
	}

	fmt.Println("\nDone.")
}
