package main

import "fmt"

func main() {
	var number [5]int
	for index := 0; index < 5; index++ {
		fmt.Println("Enter a number")
		var input int
		fmt.Scanln(&input)
		number[index] = input
	}
	for count := 0; count <= 4; count++ {
		for asterisks := 0; asterisks < number[count]; asterisks++ {
			fmt.Print("* ")
		}
		fmt.Println(" ")
	}

}
