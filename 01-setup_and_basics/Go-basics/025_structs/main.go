package main

import "fmt"

// struct groups related fields into one type

type User struct  {
	ID int
	Name string
	Email string
	Age int
}

func main() {
	u1 := User{
		ID: 1,
		Name: "Ankit",
		Email: "temp@gmail.com",
		Age: 20,
	}

	fmt.Println(u1, u1.Name)

	// mutable by default
	u1.Age = 22
	fmt.Println(u1.Age)

	// partial update
	u2 := User{
		ID: 2,
	}
	fmt.Println("partial user", u2)
}