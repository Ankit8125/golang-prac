package main

import "fmt"

func main (){
	// map [keyType] valueType

	ages := map [string] int {
		"Ankit" : 20,
		"Raj" : 22,
	}

	fmt.Println(ages, ages["Ankit"])

	// make(map[K] V)
	var scores map[string] int // It is a nil map
	fmt.Println(scores, scores["a"])

	scores = make(map[string]int)
	scores["math"] = 90

	fmt.Println(scores)

	users := map [string] int {
		"u1": 10,
		"u2": 20,
		"u3": 30,
	}

	fmt.Println(users)

	delete(users, "u2")
	fmt.Println(users)

	delete(users, "u200") // No error
	fmt.Println(users)
}