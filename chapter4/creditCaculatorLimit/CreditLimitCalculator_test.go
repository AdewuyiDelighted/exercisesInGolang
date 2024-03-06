package main

import (
	"strconv"
	"testing"
)

func TestCreditLimitCalculator(t *testing.T) {
	creditLimit := CreditLimit{
		strconv.Itoa(21345),
		200,
		400,
		500,
		500,
	}
	creditLimit.SetAccountNumber("21243")
	creditLimit.SetBeginningBalance(500)
	creditLimit.SetTotalItem(700)
	creditLimit.SetTotalCredit(900)
	creditLimit.SetAllowedCredit(1100)
	var newBalance = creditLimit.getNewBalance()
	if newBalance != 300.0 {
		t.Errorf("expected is %f  output is %f", newBalance, 300.0)

	}

}
