package main

import "fmt"

func main() {
	var value = 0
	for index := 1; index <= 100; index++ {
		value += index
		fmt.Printf("%d%10d\n", index, value)
	}
}
