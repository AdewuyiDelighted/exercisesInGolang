package main

import "fmt"

func main() {
	for row := 0; row < 10; row++ {
		for colum := 0; colum < row; colum++ {
			fmt.Print("*")
		}
		for colum := row; colum < 10; colum++ {
			fmt.Print(" ")
		}
		for colum := row; colum < 10; colum++ {
			fmt.Print("*")
		}
		for colum := 0; colum < row; colum++ {
			fmt.Print(" ")
		}
		for colum := 0; colum < row; colum++ {
			fmt.Print(" ")
		}
		for colum := row; colum < 10; colum++ {
			fmt.Print("*")
		}
		for colum := row; colum < 10; colum++ {
			fmt.Print(" ")
		}
		for colum := 0; colum < row; colum++ {
			fmt.Print("*")
		}
		fmt.Println("")

	}

}
