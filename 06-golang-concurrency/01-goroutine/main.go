package main

import (
	"fmt"
	"time"
)

func main () {
	/*
	Concepts:
	1. Concurency:
	- Doing multiple tasks in a way that they can overlap in time.
	- Overlap -> while Task A (Network, DB call), Task B can run.

	* Overlapping -> Concurrency
	* Executing multiple tasks at the exact same time (multiple CPU cores are involved)

	- Golang code can be concurrent without being parallel
		- I/O bound work -> Slow (because of waiting)
		- HTTP Req., DB Queries, Calling Microservices, Waiting for times/retries

	Read: Concurency vs Parallelism
	*/

	/*
	Goroutine: (Basically, concurrent execution of a function)
	Defn: A goroutine is a lightweight, independently executing function in Go, managed by the Go runtime, allowing 
	for efficient concurrency by running concurrently with other functions in the same program, making complex 
	async tasks simpler and more performant than traditional threads. 
	You start one by simply adding the go keyword before a function call, like go myFunc(), and they're cheap to create, 
	often using only a few kilobytes of stack memory, enabling thousands to run simultaneously. 
	*/

	start := time.Now()
	
	go func ()  {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("goroutine A: finish simulated API at", time.Since(start))
	}()

	go func ()  {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("goroutine B: finish simulated API at", time.Since(start))
	}()

	// Main -> does not wait after starting goroutine 
	fmt.Println("main: started two goroutines at ", time.Since(start))

	// small work - any logic
	fmt.Println("main: doing step 1 ", time.Since(start))
	time.Sleep(100*time.Millisecond)

	fmt.Println("main: doing step 2 ", time.Since(start))
	time.Sleep(100*time.Millisecond)

	fmt.Println("main: doing step 3 ", time.Since(start))

	// If the main function ends before goroutines fetches the value, we might not be able to get the result from goroutine.

	// Temporary sleep time
	time.Sleep(500 * time.Millisecond)

	fmt.Println("main: exit at ", time.Since(start))
}