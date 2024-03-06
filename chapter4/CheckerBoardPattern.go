package main

import "fmt"

func main() {
	for index := 0; index < 4; index++ {
		for inner := 0; inner < 8; inner++ {
			fmt.Print("* ")
		}
		fmt.Println(" ")
		for innerTwo := 0; innerTwo < 8; innerTwo++ {
			fmt.Print(" *")
		}
		fmt.Println(" ")
	}

}
