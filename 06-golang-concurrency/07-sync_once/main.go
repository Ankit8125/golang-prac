package main

import (
	"fmt"
	"sync"
	// "time"
)

func main () {
	var wg sync.WaitGroup
	var once sync.Once

	workers := 5

	for i := 1; i<= workers; i++ {
		wg.Add(1)

		go func (id int) {
			defer wg.Done()
			once.Do(setUp) // we are doing this only once. 
			// Note: only 1 goroutine is going to execute this setup. Until then, other goroutines wait until this "setUp" goroutine finishes.

			fmt.Println("worker" , i, " running")
		} (i)
	}

	wg.Wait()
	fmt.Println("main is now done")
}

func setUp() {
	fmt.Println("init: starting")
	// time.Sleep(600 * time.Millisecond)
	fmt.Println("init is now done")
}

/*
Output:
init: starting
init is now done
worker 5  running
worker 1  running
worker 3  running
worker 4  running
worker 2  running
main is now done
*/