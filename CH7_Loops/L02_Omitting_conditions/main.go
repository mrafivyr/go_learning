package main

func maxMessages(thresh int) int {
	numMessages := 0
	totalCost := 0
	for itr := 0; ; itr++ {
		totalCost = totalCost + 100 + 1*itr
		if totalCost > thresh {
			break
		}
		numMessages++
	}
	return numMessages
}
