package main

import (
	"fmt"
	"log"
	"strconv"
)

func main () {
	// Go doesn't use exceptions for normal failures
	// Functions -> return errors as normal return values

	// val, err := something()
	// if err != nil { handle the err }

	if err := run() ; err != nil {
		log.Fatal(err)
	}
}

func run () error { // going to return an error
	input := "-1"

	level, err := parseLevel(input)
	if err != nil {
		return err
	}

	fmt.Println("selected level", level)
	return nil
}

func parseLevel (s string) (int, error) {
	// (value, err)
	// nil error -> success
	// not nil -> failure

	// pattern
	n, err := strconv.Atoi(s) // converts string to integer
	if err != nil {
		return 0, fmt.Errorf("Level must be a number.")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("Level must be 1 and 5.")
	}

	return n, nil
}