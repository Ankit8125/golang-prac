package main

import "fmt"

// In Go, an array is a fixed-size, ordered sequence of elements of the same type.
// The size of an array is part of its type, meaning [3]int and [5]int are considered different types and cannot be resized after creation.
// In Go, slices are dynamically-sized and built on top of arrays and are used more often than raw arrays.

func main(){
	var marks[3] int // fixed and cannot grow
	marks[0] = 10
	marks[1] = 30
	marks[2] = 50

	fmt.Println(marks) // prints [10, 30, 50]

	// Array literal: another way to initialize 
	res := [5] int {2,3,4,5,6}
	fmt.Println(res, len(res))
}