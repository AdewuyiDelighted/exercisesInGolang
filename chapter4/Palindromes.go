package main

import (
	"fmt"
	"strconv"
)

func main() {
	var initialValue = 0
	var reversed string
	var thousand = 1
	var oneInput = 0

	fmt.Println("Enter five digit number")
	fmt.Scanln(&initialValue)

	for initialValue <= 999 || initialValue > 9999 {
		fmt.Println("Invalid input. Please Enter five digit number")
		fmt.Scanln(&initialValue)
	}
	for index := 0; index < 4; index++ {

		oneInput = initialValue / thousand % 10
		reversed += strconv.Itoa(oneInput)
		thousand = thousand * 10

	}
	var initialValueString = strconv.Itoa(initialValue)

	if initialValueString == reversed {
		fmt.Println(initialValue, " is a palindrome number")
	} else {
		fmt.Println(initialValue, " is not a palindrome number")
	}
}
