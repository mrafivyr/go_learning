package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	totalCost := 0
	/*
		customerCost, err := sendSMS(msgToCustomer)
		if err == nil {
			totalCost += customerCost
			spouseCost, errspouse := sendSMS(msgToSpouse)
			if errspouse == nil {
				totalCost += spouseCost
				return totalCost, nil
			} else {
				return 0, errspouse
			}
		} else {
			return 0, err
		}
	*/
	customerCost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}
	totalCost += customerCost
	spouseCost, errspouse := sendSMS(msgToSpouse)
	if errspouse != nil {
		return 0, errspouse
	}
	totalCost += spouseCost
	return totalCost, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
