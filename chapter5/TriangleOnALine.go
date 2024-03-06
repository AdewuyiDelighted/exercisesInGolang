package main

import "fmt"

func main() {
	for row := 0; row < 10; row++ {
		for column := 0; column < row; column++ {
			fmt.Print("*")

		}
		fmt.Print(" ")
		for column := row; column < 10; column++ {
			fmt.Print(" ")

		}
		fmt.Print(" ")
		for column := row; column < 10; column++ {
			fmt.Print("*")

		}
		fmt.Print(" ")
		for column := 0; column < row; column++ {
			fmt.Print(" ")
        
		}
		fmt.Print(" ")
		for column := 0; column < row; column++ {
			fmt.Print(" ")

		}
		fmt.Print(" ")
		for column := row; column < 10; column++ {
			fmt.Print("*")

		}
		fmt.Print(" ")
		for column := row; column < 10; column++ {
			fmt.Print(" ")

		}
		fmt.Print(" ")
		for column := 0; column < row; column++ {
			fmt.Print("*")

		}

		fmt.Println("")

	}

}
