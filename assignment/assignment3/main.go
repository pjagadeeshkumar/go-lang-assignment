/*
Assignment 3: Mutex for Synchronization Problem: Write a Go program that uses mutexes to synchronize access to a shared variable.
Multiple goroutines should increment the variable concurrently, and the final value should be printed.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	sharedVariable := 0 // shared variable is set as zero to show the final value
	numGoroutines := 10 // Number of goroutines to increment the shared variable

	// Increment function for each goroutine
	increment := func(goroutineNumber int) {
		mutex.Lock() // Lock the mutex to ensure exclusive access to the shared variable
		fmt.Printf("Incerasing the sharedVariable by Go Routine - %v\n", goroutineNumber)
		sharedVariable += 1 // Increment the shared variable
		fmt.Printf("Current Shared variable value: %v\n", sharedVariable)
		mutex.Unlock() // Unlock the mutex after updating the shared variable
		wg.Done()      // Notify the WaitGroup that this goroutine is done
	}

	// Launch multiple goroutines to increment the shared variable concurrently
	wg.Add(numGoroutines)
	for i := 1; i <= numGoroutines; i++ {
		go increment(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of the shared variable
	fmt.Printf("Final value of the shared variable: %d\n", sharedVariable)
}
