package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "ankit",
		Age:  22,
	}
	fmt.Println(u.Age)
	u.Birthday()
	fmt.Println("after", u.Age)
}

func (u *User) Birthday() {
	u.Age++
}