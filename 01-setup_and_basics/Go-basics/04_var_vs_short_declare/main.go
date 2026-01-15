package main

import (
	"fmt"
)

func main(){
	var city string
	city = "Delhi"

	var name = "Ankit" // inferred to string automatically. Hover on 'name' to see it.

	// := -> this is declares and assigns -> type inference
	friends := 3000

	likes, comments := 200, 88

	fmt.Println(city, name, friends, likes, comments)
}