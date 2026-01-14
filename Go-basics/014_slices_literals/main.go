package main

import "fmt"

// In Go, slices are dynamically-sized and built on top of arrays and are used more often than raw arrays.

func main(){
	// most common collection type. It is dynamic and can grow
	// Syntax: [] type {...}
	
	results := [] string {"ankit", "verma"}
	fmt.Println(results, results[0], len(results)) // we can access this by indexes

	// we can also update the index
	results[0] = "Rishi"
	fmt.Println(results)

	var nums [] int
	nums = append(nums, 10)
	nums = append(nums, 30, 40)

	fmt.Println(nums)


}