package main

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for itr := 0; itr < numMessages; itr++ {
		totalCost = totalCost + 1.0 + 0.01*float64(itr)
	}
	return totalCost
}
