package main

import "fmt"

func main() {
	var additionOfInput = 0
	fmt.Println("Enter initial input number")
	var firstInput int
	fmt.Scanln(&firstInput)
	for additionOfInput != firstInput && !(additionOfInput > firstInput) {
		fmt.Println("Enter a number")
		var integer int
		fmt.Scanln(&integer)
		additionOfInput += integer
	}

}
