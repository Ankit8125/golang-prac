package main

import "fmt"

func main() {
	score := 72

	if score >=50 {
		fmt.Println("Pass")
	} else if score <50 && score >= 0 {
		fmt.Println("Fail")
	} else {
		fmt.Println("Invalid score")
	}	
}