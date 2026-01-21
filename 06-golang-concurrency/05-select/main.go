package main

import (
	"fmt"
	"time"
)

func main() {
	// select is basically like switch for our channels.
	// Interview ques: why do we need select ? give me example.

	// it lets a goroutine wait on multiple channels operations at once
	// if timeout -> then stop (with the help of select)

	resultChannel := make(chan string)

	// worker goroutine
	go func() {
		// simulate slow work / consider you are doing a network call
		time.Sleep(40 * time.Millisecond)
		// time.Sleep(400 * time.Millisecond)

		resultChannel <- "worker: success"
	}()

	// In Go, time.After(d) returns a channel that sends the current time after a duration d has elapsed.
	// It is a standard way to implement timeouts with a select statement

	// Timeour channel
	timeoutCh := time.After(250 * time.Millisecond)

	select {
	case res := <- resultChannel:
		fmt.Println("main go result ", res)

	case <- timeoutCh:
		fmt.Println("main timeout happened. stop waiting")
	}

	fmt.Println("main work is done")
}
