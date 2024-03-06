package main

import "fmt"

func main() {
	var taxRate float32
	var citizenName string
	var earning float64
	for index := 1; index <= 3; index++ {
		fmt.Println("Enter citizen", index, "name")
		fmt.Scanln(&citizenName)

		fmt.Println("Enter", citizenName, "earnings")
		fmt.Scanln(&earning)

		if earning <= 30000 {
			taxRate = float32(earning * 0.15)
		}
		if earning > 30000 {
			taxRate = float32(earning * 0.2)
		}
		fmt.Println("The tax rate for", citizenName, " is", taxRate)
	}

}
