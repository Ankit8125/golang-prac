package main

import (
	"errors"
	"fmt"
)


func main () {

	// defer resp.body.Close()
	fmt.Println("case 1: success")
	if err := doWork(true); err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("case 2: early fail")
	if err := doWork(false); err != nil {
		fmt.Println("error: ", err)
	}
}

func doWork (success bool) error {

	// resource related
	// start message -> resource acquired
	// cleanup message -> resource released

	fmt.Println("start: resource acquired")

	// 'defer' gurantees that this function runs at the end of the function, regardless of whether it exits normally or via an early return
	// LIFO order: If you defer multiple functions, they execute in Last-In-First-Out order (reverse order of declaration)
	defer fmt.Println("end: resource released") 

	if !success {
		return errors.New("something went wrong. returning early")
	}

	fmt.Println("work is processing...")
	fmt.Println("work completed")

	return nil
}