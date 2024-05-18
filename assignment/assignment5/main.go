/*
    Pipeline Pattern Problem: Implement a pipeline pattern in Go where one goroutine generates numbers,
	another squares them, and a third prints the squared numbers.
*/

package main

import (
	"fmt"
)

// generator generates numbers from start to end and sends them to the channel
func numberGenerator(start, end int, ch chan int) {
	for i := start; i <= end; i++ {
		ch <- i
	}
	close(ch)
}

// square squares each number received from the input channel and sends the result to the output channel
func square(input <-chan int, output chan<- int) {
	for num := range input {
		output <- num * num
	}
	close(output)
}

// printer prints each number received from the input channel
func printer(input <-chan int, done chan<- bool) {
	for num := range input {
		fmt.Println(num)
	}
	done <- true // Signal that printing is done
}

func main() {
	// Create channels to connect the pipeline stages
	generatorOutput := make(chan int)
	squaredNumbersOutput := make(chan int)
	done := make(chan bool) // Channel to signal when printing is done

	// Start the generator goroutine
	go numberGenerator(1, 10, generatorOutput)

	// Start the square goroutine
	go square(generatorOutput, squaredNumbersOutput)

	// Start the printer goroutine
	go printer(squaredNumbersOutput, done)

	// Wait for printing to finish
	<-done

	// All printing is done
	fmt.Println("All goroutines have completed.")
}
