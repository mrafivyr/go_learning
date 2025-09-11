package main

func getMessageCosts(messages []string) []float64 {
	messagesCost := make([]float64, len(messages))
	for index, message := range messages {
		messagesCost[index] = float64(len(message)) * 0.01
	}
	return messagesCost
}

/*
The built-in len() function returns the number of bytes, not the number of characters (runes), in a string.
This can lead to incorrect cost calculations if your application needs to handle non-ASCII characters, which can be multiple bytes long.

import (
	"unicode/utf8"
)

func getMessageCostsUnicode(messages []string) []float64 {
	messagesCost := make([]float64, len(messages))
	for index, message := range messages {
		messagesCost[index] = float64(utf8.RuneCountInString(message)) * 0.01
	}
	return messagesCost
}

*/
