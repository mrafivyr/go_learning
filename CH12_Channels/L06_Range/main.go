package main

func concurrentFib(n int) []int {
	intSlice := make([]int, 0, n)
	ch := make(chan int, n)
	go fibonacci(n, ch)
	for val := range ch {
		intSlice = append(intSlice, val)
	}
	return intSlice
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
