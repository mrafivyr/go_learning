package main

func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}
