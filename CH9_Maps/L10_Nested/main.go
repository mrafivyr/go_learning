package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int, len(names))
	for _, name := range names {
		if len(name) == 0 {
			continue // Skip empty strings
		}
		runes := []rune(name)
		firstRune := runes[0] // Get the first rune (assuming ASCII for simplicity or iterating on runes for unicode safety)

		// Lazy-initialize the inner map if it doesn't exist.
		if nameCounts[firstRune] == nil {
			nameCounts[firstRune] = make(map[string]int)
		}
		nameCounts[firstRune][name]++
	}
	return nameCounts
}

/*
func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int, len(names))
	for _, name := range names {
		runes := []rune(name)
		if nameMap, exists := nameCounts[runes[0]]; exists {
			nameMap[name]++
		} else {
			nameCounts[runes[0]] = make(map[string]int)
			nameCounts[runes[0]][name] = 1
		}
	}
	return nameCounts
}

*/
