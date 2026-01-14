package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName := "Ankit"
	lastName := "Verma"

	fullName := firstName + " " + lastName

	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))
}
