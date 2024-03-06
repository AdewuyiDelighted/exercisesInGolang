package main

import "fmt"

func main() {
	var passes = 0
	var failure = 0
	var result int

	for student := 1; student <= 10; student++ {
		fmt.Println("Enter result (1 = pass,2 = fail)")
		fmt.Scanln(&result)

		if result == 1 {
			passes += 1
		} else {
			failure += 1
		}
		for result != 1 && result != 2 {
			fmt.Println("Invalid input Enter result (1 = pass,2 = fail)")
			fmt.Scanln(&result)

		}
	}
	if passes > 8 {
		fmt.Println("Bonus to instructor")
	}
	fmt.Printf("Passed %d\nFailed %d\n", passes, failure)

}
