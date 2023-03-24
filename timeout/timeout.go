/*
	This pattern is when we want to do something after a set interval of
	time. Using time.After() we get a channel input at the set interval and
	can capture the data at that moment to do something. 
	Here in the example we use that moment to quit.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Infinitely produces random numbers
func producer() <-chan int {
	channel := make(chan int)

	go func() {
		for {
			channel <- rand.Intn(100)
		}
	}()
	return channel
}

// This keeps consuming data untill the timeout duration is reached
// and then it enters the select-case to return
func consumer(input <-chan int) {
	// After the timeout duration, it returns a one time channel input
	timeout := time.After(time.Duration(50000 * time.Nanosecond))

	for {
		select {
		case data := <-input:
			fmt.Printf("Data received: %d\n", data)
		case <-timeout:
			fmt.Printf("Time out for receiving any more data. Quitting...\n")
			return
		}
	}
}

func main() {
	producerChannel := producer()
	consumer(producerChannel)
}