package main

/*
func countReports(numSentCh <-chan int) int {
	totalReports := 0
	for {
		if chValue, ok := <-numSentCh; !ok {
			break
		} else {
			totalReports += chValue
		}
	}
	return totalReports
}
*/

// The idiomatic and more readable approach
func countReports(numSentCh <-chan int) int {
	totalReports := 0
	for chValue := range numSentCh {
		totalReports += chValue
	}
	return totalReports
}

// don't touch below this line

func sendReports(numBatches int, ch chan<- int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
