package main

import (
	"strings"
)

/*
func countDistinctWords(messages []string) int {
	if len(messages) == 0 || messages[0] == "" {
		return 0
	}
	wordMap := make(map[string]struct{})
	for _, msg := range messages {
		msgSlice := strings.Split(msg, " ")
		for _, stringData := range msgSlice {
			wordMap[strings.ToLower(stringData)] = struct{}{}
		}
	}
	return len(wordMap)
}
*/

func countDistinctWords(messages []string) int {
	wordMap := make(map[string]struct{})
	for _, msg := range messages {
		// Use strings.Fields to handle multiple spaces and trim whitespace.
		words := strings.Fields(msg)
		for _, word := range words {
			// strings.Fields handles empty strings, so no extra check is needed.
			wordMap[strings.ToLower(word)] = struct{}{}
		}
	}
	return len(wordMap)
}
