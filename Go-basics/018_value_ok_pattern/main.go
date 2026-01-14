package main

import "fmt"

func main(){
	points := map [string] int {
		"a": 10,
		"b": 0, // Valid value
	}

	fmt.Println(points["a"])
	fmt.Println(points["b"])
	fmt.Println(points["c"])

	valB, okB := points["b"]
	fmt.Println(valB, okB)

	valC, okC := points["c"]
	fmt.Println(valC, okC)

	if val, ok := points["c"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("c is not present in map")
	}

	prices := map [string] int {
		"ab": 10,
		"ds": 100,
	}

	total := 0
	for item, price := range prices{
		fmt.Println(item, price)
		total += price
	}

	fmt.Println(total)

	// If you want only keys
	for item := range prices{
		fmt.Println(item)
	}
}
