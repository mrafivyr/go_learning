package main

func adder() func(int) int {
	// ?
	sum := 0
	return func(an int) int {
		sum += an
		return sum
	}
}
