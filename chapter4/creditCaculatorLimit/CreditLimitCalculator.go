package main

type CreditLimit struct {
	accountNumber    string
	beginningBalance float64
	totalItem        float64
	totalCredit      float64
	allowedCredit    float64
}

func (creditLimit *CreditLimit) AccountNumber() string {
	return creditLimit.accountNumber
}

func (creditLimit *CreditLimit) SetAccountNumber(accountNumber string) {
	creditLimit.accountNumber = accountNumber
}

func (creditLimit *CreditLimit) BeginningBalance() float64 {
	return creditLimit.beginningBalance
}

func (creditLimit *CreditLimit) SetBeginningBalance(beginningBalance float64) {
	creditLimit.beginningBalance = beginningBalance
}

func (creditLimit *CreditLimit) TotalItem() float64 {
	return creditLimit.totalItem
}

func (creditLimit *CreditLimit) SetTotalItem(totalItem float64) {
	creditLimit.totalItem = totalItem
}

func (creditLimit *CreditLimit) TotalCredit() float64 {
	return creditLimit.totalCredit
}

func (creditLimit *CreditLimit) SetTotalCredit(totalCredit float64) {
	creditLimit.totalCredit = totalCredit
}

func (creditLimit *CreditLimit) AllowedCredit() float64 {
	return creditLimit.allowedCredit
}

func (creditLimit *CreditLimit) SetAllowedCredit(allowedCredit float64) {
	creditLimit.allowedCredit = allowedCredit
}
func (creditLimit *CreditLimit) getNewBalance() float64 {
	return creditLimit.beginningBalance + creditLimit.totalItem - creditLimit.totalCredit
}
