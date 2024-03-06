package main

import "fmt"

func main() {
	var square, cube, cubic int
	for number := 1; number <= 5; number++ {
		square = number * number
		cube = square * number
		cubic = cube * square
		fmt.Printf("%5d %5d %5d %5d\n", number, square, cube, cubic)

	}

}
