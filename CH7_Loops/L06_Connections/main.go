package main

// func countConnections(groupSize int) int {
// 	totalCOnnections := 0
// 	for i := 1; i <= groupSize; i++ {
// 		for j := i; j <= groupSize; j++ {
// 			if i != j {
// 				totalCOnnections++
// 			}
// 		}
// 	}
// 	return totalCOnnections
// }

func countConnections(groupSize int) int {
	totalConnections := 0
	for i := 1; i < groupSize; i++ {
		totalConnections += i
	}
	return totalConnections
}
