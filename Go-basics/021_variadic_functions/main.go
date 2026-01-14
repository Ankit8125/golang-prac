package main

import "fmt"

func sumAll(nums ...int) int { // variadic function = can accept any number of arguments
	total := 0

	for _, curr := range nums {
		total += curr
	}

	return total
}

func main () {
	fmt.Println(sumAll(1,2,3,4,5,6,7))

	values := [] int {10, 20}
	fmt.Println(sumAll(values...))

	res := func (n int) int {
		return n * 2
	}

	fmt.Println(res(2))

	// IIFE
	res1 := func (a int, b int) int {
		return a + b
	}(2, 9)

	fmt.Println(res1)
}