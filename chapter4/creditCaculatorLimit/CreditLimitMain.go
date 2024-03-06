package main

import "fmt"

func main() {
	creditLimit := CreditLimit{
		"212",
		45,
		100,
		200,
		400,
	}
	var accountNumber string
	var beginning float64
	var charges float64
	var credit float64
	var creditLimits float64

	fmt.Println("Enter user account number")
	fmt.Scanln(&accountNumber)
	creditLimit.SetAccountNumber(accountNumber)

	fmt.Println("Enter the beginning balance ", accountNumber)
	fmt.Scanln(&beginning)
	creditLimit.SetBeginningBalance(beginning)

	fmt.Println("Enter the total charge")
	fmt.Scanln(&charges)
	creditLimit.SetTotalItem(charges)

	fmt.Println("Enter the total credit")
	fmt.Scanln(&credit)
	creditLimit.SetTotalCredit(credit)

	fmt.Println("Enter the credit limit of the user with account number ", accountNumber)
	fmt.Scanln(&creditLimits)
	creditLimit.SetAllowedCredit(creditLimits)

	var newBalance = creditLimit.getNewBalance()
	fmt.Println("The user with account number ", accountNumber, "new balance is", newBalance)

	if creditLimits > newBalance {
		fmt.Println("Credit limit exceeded")
	}

}
