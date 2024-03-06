package main

import "fmt"

func main() {

	for index := 0; index < 10; index++ {
		for count := 0; count < index; count++ {
			fmt.Print("* ")
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")

	for index := 0; index < 10; index++ {
		for count := index; count < 10; count++ {
			fmt.Print("* ")
		}
		fmt.Println(" ")
	}

	fmt.Println(" ")

	for index := 0; index < 10; index++ {
		for count := 0; count < index; count++ {
			fmt.Print(" ")
		}
		fmt.Print("")

		for count := index; count < 10; count++ {
			fmt.Print("*")
		}
		fmt.Println(" ")

	}
	fmt.Println(" ")
	for index := 0; index < 10; index++ {
		for count := index; count < 10; count++ {
			fmt.Print(" ")
		}
		fmt.Print("")

		for count := 0; count < index; count++ {
			fmt.Print("*")
		}
		fmt.Println(" ")

	}
}
