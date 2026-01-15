package main

import (
	"errors"
	"fmt"
)


func main () {

	// defer resp.body.Close()
}

func doWork (success bool) error {

	// resource related
	// start message -> resource acquired
	// cleanup message -> resource released

	fmt.Println("start: resource acquired")

	// 'defer' gurantees that this function runs at the end of the function
	defer fmt.Println("end: resource released") 

	if !success {
		return errors.New("something went wrong. returning early")
	}

	fmt.Println("work is processing...")
	fmt.Println("work completed")

	return nil
}