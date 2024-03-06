package main

import "fmt"

func main() {
	fmt.Println("Enter the length of the base of a triangle")
	var length int
	fmt.Scanln(&length)
	if length < 1 || length > 10 {
		fmt.Println("Invalid input")
	} else {
		for base := 1; base <= length; base++ {
			for count := 0; count < base; count++ {
				fmt.Print("* ")

			}
			fmt.Println(" ")

		}
	}

}
