package main

import (
	"errors"
	"fmt"
	"time"
)

// With Generic logic
func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	cost := newItem.GetCost()
	if cost > balance {
		return *new([]T), 0.0, errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	balance = balance - cost
	return oldItems, balance, nil
}

/*
// Without Generic logic
func chargeForLineItem(newItem lineItem, oldItems []lineItem, balance float64) ([]lineItem, float64, error) {
	cost := newItem.GetCost()
	if cost > balance {
		var tempData []lineItem
		return tempData, 0.0, errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	balance = balance - cost
	return oldItems, balance, nil
}
*/
// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}
