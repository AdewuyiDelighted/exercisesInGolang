package main

import "fmt"

func main() {
	fmt.Println("Enter first number")
	var firstNumber int
	fmt.Scanln(&firstNumber)
	fmt.Println("Enter second number")
	var secondNumber int
	fmt.Scanln(&secondNumber)

	if firstNumber == secondNumber {
		fmt.Println("Result", 0)
	}
	if firstNumber > secondNumber {
		fmt.Println("Result", 1)
	}
	if secondNumber > firstNumber {
		fmt.Println("Result", -1)
	}

}
