package main

import "fmt"

func main() {
	// store the memory address of any value

	// &ankit -> address of "ankit" (makes a pointer i.e. it points to the memory location of "ankit" variable)
	// *p -> dereference (go to that address and read/write)
	// I/W QUES: why do we need it ? => ANS: to change the value of something inside a function without returning it.

	score := 10
	fmt.Println("before: ", score)

	addScore(&score)

	fmt.Println("after: ", score)
}	

func addScore(score *int){
	*score = 5;
}	