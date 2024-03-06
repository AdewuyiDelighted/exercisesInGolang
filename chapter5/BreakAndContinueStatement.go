package main

import "fmt"

func main() {

	fmt.Println("Enter the length of number you like to loop over")
	var length int
	fmt.Scanln(&length)
	fmt.Println("Enter the number you would like to break at")
	var breakValue int
	fmt.Scanln(&length)

	for index := 1; index <= length; index++ {
		if index > breakValue-1 {
			index = index + length + 1
		} else {
			fmt.Println(index)
		}
	}

}
