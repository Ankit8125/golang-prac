package main

import "fmt"

func main(){
	views := [] int {10, 20, 30}

	// for range
	total := 0
	for i, v := range views{
		fmt.Println("Day ", i, " Views ", v)
		total += v
	}

	fmt.Println(total)
}