package main

import (
	"fmt"
	"time"
)

func main () {
	// waitgroups help you to wait

	// It's like a pipe which is going to exchange values between goroutines
	
	// One goroutine sends:  ch <- value
	// another receives:  value := <- ch

	// make(chan T) -> It is going to create a unbuffered channel
	
	type User struct {
		ID int
		Name string
	}

	ch := make (chan User) // this "ch" carries the User values

	// worker goroutine
	go func () {

		// simulate slow work
		time.Sleep(200*time.Millisecond)
		
		// Send: blocks until main receives
		// Unbuffered channel, send + receives is like a handshake, 
		ch <- User { ID: 100, Name: "Ankit" } // We are sending the User values into the channel
	
	} ()

	fmt.Println("Main is now waiting to receive user...")

	v := <- ch

	fmt.Println("Got user", v, v.ID, v.Name)
}