package main

import "fmt"

func main() {
	for day := 1; day <= 12; day++ {
		switch day {
		case 1:
			fmt.Println("On The first Day Of Christmas My true love sent to me")
			break
		case 2:
			fmt.Println("On The second Day Of Christmas My true love sent to me")
			break
		case 3:
			fmt.Println("On The third Day Of Christmas My true love sent to me")
			break
		case 4:
			fmt.Println("On The fourth Day Of Christmas My true love sent to me")
			break
		case 5:
			fmt.Println("On The fifth Day Of Christmas My true love sent to me")
			break
		case 6:
			fmt.Println("On The sixth Day Of Christmas My true love sent to me")
			break
		case 7:
			fmt.Println("On The seventh Day Of Christmas My true love sent to me")
			break
		case 8:
			fmt.Println("On The eighth Day Of Christmas My true love sent to me")
			break
		case 9:
			fmt.Println("On The ninth Day Of Christmas My true love sent to me")
			break
		case 10:
			fmt.Println("On The tenth Day Of Christmas My true love sent to me")
			break
		case 11:
			fmt.Println("On The eleventh Day Of Christmas My true love sent to me")
			break
		case 12:
			fmt.Println("On The twelfth Day Of Christmas My true love sent to me")
			break
		}

		switch day {
		case 12:
			fmt.Println("Twelve drummers drumming")
			fallthrough
		case 11:
			fmt.Println("Eleven piper piping")
			fallthrough
		case 10:
			fmt.Println("Ten lords a-leaping")
			fallthrough
		case 9:
			fmt.Println("Nine ladies dancing")
			fallthrough
		case 8:
			fmt.Println("Eight maids a-milking")
			fallthrough
		case 7:
			fmt.Println("Seven swan a-swimming")
			fallthrough
		case 6:
			fmt.Println("Six  geese a-laying")
			fallthrough
		case 5:
			fmt.Println("Five golden ring")
			fallthrough
		case 4:
			fmt.Println("Four calling birds")
			fallthrough
		case 3:
			fmt.Println("Three French hens")
			fallthrough
		case 2:
			fmt.Println("Two turtle dove")
			fallthrough
		case 1:
			fmt.Println("And a partridge in pear tree")

		}

	}

}
