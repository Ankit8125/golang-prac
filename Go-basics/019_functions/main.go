package main

import "fmt"

func add(a int, b int) int {
	return a+b
}

// If we want to return multiple values
func SumAndProduct(a int, b int) (int, int) {
	sum := a + b
	prod := a * b

	return sum, prod
}

func main() {
	sum1 := add(10, 20)
	fmt.Println(sum1)	

	s, p := SumAndProduct(10, 92) // we can use "_" if we do not want any value
	s1, _ := SumAndProduct(20, 20)
	fmt.Println(s, p, s1)
}