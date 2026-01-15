package main

import (
	"fmt"
	"go-modules/internal/greet"
)

func main () {
	msg1 := greet.Hello("Ankit")

	fmt.Println(msg1)
}