package main

import "fmt"

func fizzbuzz() {
	for itr := 1; itr <= 100; itr++ {
		if itr%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if itr%3 == 0 {
			fmt.Println("fizz")
		} else if itr%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(itr)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
