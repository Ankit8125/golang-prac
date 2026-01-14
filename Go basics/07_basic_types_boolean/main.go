package main

import (
	"fmt"
)

func main(){
	isLoggedIn := true // Inferred as boolean type
	isAdmin := true

	// AND &&
	canOpen := isLoggedIn && isAdmin // similarly we have OR: ||
	fmt.Println("Is user logged in?", canOpen)	
}
