package main

import (
	"fmt"
	"time"
)

// This function will run concurrently in a goroutine.
// It sends a message to the channel `c`.
func worker(c chan string) {
	fmt.Println("Worker started...")
	time.Sleep(time.Second * 2) // Simulate some work
	fmt.Println("Worker finished work.")
	c <- "Hello from the worker goroutine!" // Send a value into the channel.
}

func goroutines() {
	// Channels are the pipes that connect concurrent goroutines.
	// You can send values into channels from one goroutine
	// and receive those values into another.
	messages := make(chan string)

	// To start a new goroutine, use the `go` keyword followed by a function call.
	// This will execute the `worker` function concurrently with the `main` function.
	go worker(messages)

	fmt.Println("Main function is waiting for a message...")

	// This line blocks until a message is received from the channel.
	// This is how we synchronize the main function with the worker goroutine.
	msg := <-messages

	fmt.Println("Main function received a message:")
	fmt.Println(msg)
}
