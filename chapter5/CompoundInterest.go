package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("%s%10s%13s%15s%16s%15s%14s\n", "Year", "5%", "6%", "7%", "8%", "9%", "10%")
	var principal = 1000 + 0.1
	for year := 1; year <= 10; year++ {
		for interest := 0.05; interest <= 0.1; interest += 0.01 {
			var rate = principal * math.Pow(1.0+interest, float64(year))
			fmt.Printf("%d%14.1f", year, rate)
		}
		fmt.Println()
	}

}
