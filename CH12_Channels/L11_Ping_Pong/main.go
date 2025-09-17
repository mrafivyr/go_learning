/*
package main

import (
	"fmt"
	"time"
)

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	go ponger(pings, pongs)
	go pinger(pings, numPings)
	go func() {
		i := 0
		for range pongs {
			fmt.Println("got pong", i)
			i++
		}
		fmt.Println("pongs done")
	}()
}

// don't touch below this line

func pinger(pings chan struct{}, numPings int) {
	sleepTime := 50 * time.Millisecond
	for i := 0; i < numPings; i++ {
		fmt.Printf("sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
}
*/

// Logic to wait main Goroutine until all 3 child Goroutine's are finished
package main

import (
	"fmt"
	"sync"
	"time"
)

func pingPong(numPings int) {
	var wg sync.WaitGroup

	pings := make(chan struct{})
	pongs := make(chan struct{})

	// Add 3 goroutines to the wait group
	wg.Add(3)

	// Ponger goroutine
	go func() {
		defer wg.Done()
		ponger(pings, pongs)
	}()

	// Pinger goroutine
	go func() {
		defer wg.Done()
		pinger(pings, numPings)
	}()

	// Pong-reader goroutine
	go func() {
		defer wg.Done()
		i := 0
		for range pongs {
			fmt.Println("got pong", i)
			i++
		}
		fmt.Println("pongs done")
	}()

	// Wait for all goroutines to finish before returning
	wg.Wait()
}

// don't touch below this line

func pinger(pings chan struct{}, numPings int) {
	sleepTime := 50 * time.Millisecond
	for i := 0; i < numPings; i++ {
		fmt.Printf("sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
}
