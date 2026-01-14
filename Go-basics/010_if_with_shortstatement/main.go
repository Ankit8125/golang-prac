package main

import "fmt"

func main() {
	items := 3
	pricePerItem := 40

	if total := items * pricePerItem; total > 100 {
		fmt.Println("Eligible for shopping")
	} else {
		fmt.Println("Not eligible")
	}
}