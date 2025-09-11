package main

/*
func indexOfFirstBadWord(msg []string, badWords []string) int {
	indexFound := -1
	for index, message := range msg {
		indexFound = slices.Index(badWords, message)
		if indexFound >= 0 {
			return index
		}
	}
	return indexFound
}
*/
// modified with contains
/*
func indexOfFirstBadWord(msg []string, badWords []string) int {
	for index, message := range msg {
		if slices.Contains(badWords, message) {
			return index
		}
	}
	return -1
}
*/

// Using map .. map[string]struct{} instead of map[string]bool because of size and clarity
func indexOfFirstBadWord(msg []string, badWords []string) int {
	// Create a set (map) of bad words for efficient lookups.
	badWordSet := make(map[string]struct{}, len(badWords))
	for _, word := range badWords {
		badWordSet[word] = struct{}{}
	}

	// Loop through the messages and check for existence in the set.
	for index, message := range msg {
		if _, exists := badWordSet[message]; exists {
			return index
		}
	}

	// Return -1 if no bad words are found.
	return -1
}
