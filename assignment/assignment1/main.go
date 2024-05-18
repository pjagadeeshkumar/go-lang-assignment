/*
Assignment 1: Goroutine with Channel Problem: Write a Go program that calculates the
sum of numbers from 1 to N concurrently using goroutines and channels.
The program should take the value of N as input from the user.
*/
package main

import (
	"fmt"
)

// worker function that calculates the sum of numbers from 1 to N
func sumToN(N int, ch chan int) {
	sum := 0
	for i := 1; i <= N; i++ {
		sum += i
	}
	ch <- sum // send the result to the channel
	close(ch) // close the channel after sending the result
}

func main() {
	var num uint

	// Create a channel to receive the sum from each goroutine
	ch := make(chan int)

	fmt.Print("Enter the value of Num (range): ")
	fmt.Scan(&num)

	go sumToN(int(num), ch)

	totalSum := <-ch

	fmt.Printf("\nThe sum of numbers from 1 to %d is: %d\n", num, totalSum)

}
