package main

import "fmt"

func divide (a int, b int) (ankit int, verma int) { // named return values
	ankit = a+b
	verma = a-b

	return // aka "Naked Return". This is going to return values in the named variables (ankit, verma) above.
}

func main() {
	q, r := divide(10, 5)
	fmt.Println(q, r)
}