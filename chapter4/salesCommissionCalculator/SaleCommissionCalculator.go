package main

import "fmt"

func main() {
	var noOfItems int
	var priceOfItems float64
	var itemTotalPrice float64

	fmt.Println("Enter the number of items sold by this salesperson")
	fmt.Scanln(&noOfItems)

	for item := 0; item < noOfItems; item++ {
		fmt.Println("Enter the price of item ", item+1)
		fmt.Scanln(&priceOfItems)

		itemTotalPrice += priceOfItems
	}
	var commission = itemTotalPrice * 0.09
	var earnings = commission + 200
	fmt.Printf("%s%.1f", "The salesperson earning is ", earnings)
}
