package main

import (
	"context"
	"fmt"
	"time"
)

func main () {
	// context.withTimeout(parent, timeout) returns a derived context that is automatically cancelled after a specific duration.

	// context.Background() is the root of all context trees in Go. It is an empty context that is never cancelled, has no values and has no deadline. 
	ctx, cancel := context.WithTimeout(context.Background(), 450 * time.Millisecond)
	defer cancel() // what if i do not release the resources ? that is i do not write defer in my curent vscode. when i close it, then will it be lost or.. ?

	go slowWork(ctx)

	// main waits until context ends
	<- ctx.Done()

	fmt.Println("main context ended: ", ctx.Err())
	fmt.Println("main exit")
}

func slowWork(ctx context.Context){
	select {
	case <- time.After(700*time.Millisecond):
		fmt.Println("slow work is done")
		return
	
	case <- ctx.Done():
		fmt.Println("slow work stopped:", ctx.Err())
		return
	}
}