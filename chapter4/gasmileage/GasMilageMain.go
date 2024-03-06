package main

import "fmt"

func main() {
	gasMileage := GasMileage{
		0,
		0,
	}
	var count = 0.0
	var noOfTrips = 0.0
	var milesPerGallon = 0.0
	var totalMiles = 0.0
	var totalGallon = 0.0
	var totalMilesPerGallon = 0.0

	for noOfTrips != -1 {
		fmt.Println("Enter the miles for trip ", count+1)
		var miles float64
		fmt.Scanln(&miles)
		gasMileage.setMile(miles)
		totalMiles += miles

		fmt.Println("Enter the gallons used for trip ", count+1)
		var gallon float64
		fmt.Scanln(&gallon)
		gasMileage.setGallon(gallon)
		totalGallon += gallon

		milesPerGallon = float64(gasMileage.calculateMilesPerGallons())

		fmt.Println("The miles per gallon for trip ", count+1, " is", milesPerGallon)
		count += 1
		fmt.Println("Enter -1 to exit")
		fmt.Scanln(&noOfTrips)
	}
	totalMilesPerGallon = totalMiles / totalGallon

	fmt.Printf("%s%.1f", "The total miles per gallon for all trips  is ", totalMilesPerGallon)

}
