package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1 : // "1" is the value
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Weekday")
	}
}