package main

// imports brings external packages into the file that you are working, where you actually need it.
import (
	"fmt" // standard package for "formatting" input or output
	"math"
)

func main(){
	// packageName.functionName -> call a function from a package
	fmt.Println("sqrt(25)", math.Sqrt(25))
}