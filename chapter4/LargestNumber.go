package main

import "fmt"

func main() {
	var largest = 0
	var number int
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Enter input", counter)
		fmt.Scan(&number)
		if number > largest {
			largest = number
		}
	}
	fmt.Println("The largest number is ", largest)

}
