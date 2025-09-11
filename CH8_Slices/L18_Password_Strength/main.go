package main

import (
	"unicode"
)

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	var upperCaseChar, digitChar bool
	for _, char := range password {
		if unicode.IsUpper(char) {
			upperCaseChar = true
		}
		if unicode.IsDigit(char) {
			digitChar = true
		}

		if upperCaseChar && digitChar {
			return true
		}
	}

	return false
}
