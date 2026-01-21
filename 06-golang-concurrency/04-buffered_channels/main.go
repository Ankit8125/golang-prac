package main

import (
	"fmt"
	"time"
)

func main () {
	// unbuffered -> is like a handshake.
	// Sender blocks until receiver is ready and vice versa. (Both send and receive at the middle)

	// Buffered channel -> capacity -> buffer size
	// for ex: make(chan int, 3)

	jobs := make(chan string, 2) // 2 means; channel can store 2 jobs without a receiver

	// Producer
	go func () {
		fmt.Println("producer: sending job-1")
		jobs <- "job-1" // goes into the buffer (as buffer has limited space now)

		fmt.Println("producer: sending job-2")
		jobs <- "job2" // buffer still has space
		// buffer gets full. So now sender is blocked UNTIL the consumer is going to receives atleast 1 job, so that we get atleast 1 free space.
		
		fmt.Println("producer: sending job-3 but this will wait until consumer reads")
		jobs <- "job3"

		fmt.Println("producer: sent all jobs")
		
		close(jobs) // closes the sender -> no more jobs, all done
 	} ()

	// Consumer
	for job := range jobs {
		fmt.Println("Consumer got ", job)
		time.Sleep(300*time.Millisecond)
		fmt.Println("consumer finished ", job)
	}

	fmt.Println("Main: all jobs completed")
}