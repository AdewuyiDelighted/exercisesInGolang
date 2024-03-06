package main

import (
	"fmt"
	"strconv"
)

func main() {
	var binaryNumber = 0
	var oneInput = 0
	//var power = 0
	var count = 1
	var strValueOfBinaryNumber string
	var length = 0
	var decimalValue = 0

	fmt.Println("Enter a binary number")
	fmt.Scanln(&binaryNumber)
	strValueOfBinaryNumber = strconv.Itoa(binaryNumber)

	length = len(strValueOfBinaryNumber)

	for index := 0; index < length; index++ {
		oneInput = binaryNumber / count % 10

		decimalValue += oneInput * findSquareRoot(2, index)
		count *= 10
	}
	fmt.Println(decimalValue)
}

func findSquareRoot(input, raisePower int) int {
	var sqrRoot = 1
	for index := 0; index < raisePower; index++ {
		sqrRoot *= input
	}
	return sqrRoot
}
