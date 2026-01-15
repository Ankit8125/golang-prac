package main

import "fmt"

func main(){

	// make([] datatype, length, capacity)
	scores := make([] int, 0, 2)
	fmt.Println(len(scores), cap(scores), scores)

	scores = append(scores, 100, 200)
	fmt.Println(len(scores), cap(scores), scores)

	// Incase, we are exceeding capacity, Go grows the backing array (ususally doubles until a particular size)
	scores = append(scores, 22)
	fmt.Println(len(scores), cap(scores), scores)

	todos := [] string {"gym", "eat"}
	more := [] string {"code"}

	todos = append(todos, more...)
	fmt.Println(todos)
}