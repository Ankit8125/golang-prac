package main

import "fmt"

type User struct {
	Name string
	Age int	
}

func main () {

	u1 := User{
		Name: "Ankit",
		Age: 22,
	}

	fmt.Println(u1.Intro())

}

// "val receiver" means this method receives a copy of the struct (User)
func (u User) Intro() string { // Any changes inside the method don't affect the original u1.
	return fmt.Sprintf("Hi, I am %s", u.Name)
}

