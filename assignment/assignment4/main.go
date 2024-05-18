/*
	Goroutine with Timeout Problem: Write a Go program that performs a task in a goroutine and waits for it to finish.
	However, if the task takes more than 3 seconds, the program should print a timeout message and exit.
*/

package main

import (
	"fmt"
	"time"
)

func task() {
	// Simulate some long-running task
	time.Sleep(5 * time.Second)
	fmt.Println("Task completed")
}

func main() {
	// Create a channel to signal completion of the task
	done := make(chan bool)

	// Start the task in a goroutine
	go func() {
		task()
		done <- true // Signal that the task is completed
	}()

	// Wait for either the task to complete or for timeout
	select {
	case <-done:
		fmt.Println("Task completed within 3 seconds.")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout: Task took more than 3 seconds.")
	}
}
