package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(cust *customer, trans transaction) error {
	if trans.transactionType != transactionDeposit && trans.transactionType != transactionWithdrawal {
		return errors.New("unknown transaction type")
	}

	if trans.transactionType == transactionWithdrawal {
		if trans.amount > cust.balance {
			return errors.New("insufficient funds")
		}
		cust.balance = cust.balance - trans.amount
	}

	if trans.transactionType == transactionDeposit {
		cust.balance = cust.balance + trans.amount
	}

	return nil
}
