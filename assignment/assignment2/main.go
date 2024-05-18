/*
Assignment 2: Producer-Consumer with Channel Problem: Implement the producer-consumer problem using goroutines and channels.
The producer should generate numbers from 1 to 100 and send them to a channel, and the consumer should print those numbers.
*/

package main

import "fmt"

// producer program to generate 1 to 100 numbers and feed into channel
func producer(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch) // Close the channel after sending all numbers
}

// consumer program to print the numbers in channel
func consumer(ch chan int, done chan bool) {
	for num := range ch {
		fmt.Println(num)
	}
	done <- true // Signal that the consumer is done
}

func main() {
	// creating a channel to store and process numbers
	ch := make(chan int)

	// creating a channel to signal consumer is done
	done := make(chan bool)

	// invoke producer as go routine
	go producer(ch)

	// invoke consumer to print numbers while signalling when done
	go consumer(ch, done)

	<-done

}
