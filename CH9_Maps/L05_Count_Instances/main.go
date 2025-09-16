package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, exists := validUsers[user]; exists {
			validUsers[user]++
		}
	}
}
