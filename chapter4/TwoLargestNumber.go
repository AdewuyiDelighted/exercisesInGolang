package main

import "fmt"

func main() {
	fmt.Println("Enter a number")
	var input int
	fmt.Scanln(&input)
	var firstLargest = input
	var secondLargest = 0

	for index := 0; index < 9; index++ {
		fmt.Println("Enter a number")
		fmt.Scanln(&input)
		if input > firstLargest {
			secondLargest = firstLargest
			firstLargest = input
		}
		if input > secondLargest && input < firstLargest {
			secondLargest = input
		}
	}
	fmt.Println("The first highest ", firstLargest)
	fmt.Println("The second highest ", secondLargest)

}
