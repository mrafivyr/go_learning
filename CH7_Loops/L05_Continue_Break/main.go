package main

import (
	"fmt"
)

func printPrimes(max int) {
	if max >= 2 {
		fmt.Println(2)
	}
	for itr := 3; itr <= max; itr += 2 { // Start with 3 and step by 2, skipping all even numbers
		isPrime := true
		for inner_itr := 3; inner_itr*inner_itr <= itr; inner_itr += 2 { // Check divisibility only up to the square root, and only with odd numbers
			if itr%inner_itr == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(itr)
		}
	}
}

// func printPrimes(max int) {
// 	if max < 2 {
// 		return
// 	}

// 	isPrime := make([]bool, max+1)
// 	for i := 2; i <= max; i++ {
// 		isPrime[i] = true
// 	}

// 	for p := 2; p*p <= max; p++ {
// 		if isPrime[p] {
// 			for i := p * p; i <= max; i += p {
// 				isPrime[i] = false
// 			}
// 		}
// 	}

// 	for i := 2; i <= max; i++ {
// 		if isPrime[i] {
// 			fmt.Println(i)
// 		}
// 	}
// }

// // don't edit below this line

// func test(max int) {
// 	fmt.Printf("Primes up to %v:\n", max)
// 	printPrimes(max)
// 	fmt.Println("===============================================================")
// }

//	func main() {
//		test(10)
//		test(20)
//		test(30)
//		test(3000000)
//	}

func printPrimesUpTo(max int, lowRange, highRange int) {
	if max < 2 {
		return
	}

	isPrime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= max; p++ {
		if isPrime[p] {
			for i := p * p; i <= max; i += p {
				isPrime[i] = false
			}
		}
	}

	for i := lowRange; i <= highRange; i++ {
		if i >= 2 && isPrime[i] {
			fmt.Println(i)
		}
	}
}

func main() {
	low := 2985100
	high := 2985200
	printPrimesUpTo(high, low, high)
	// printPrimesUpTo(100, 90, 100)
	// printPrimes(100)
}
