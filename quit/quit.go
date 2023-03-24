/*
	Quit pattern basically is a way to quit some work being
	done using channels by using a different channel that can be
	used to push data and then when data is received in that channel,
	it is time to quit.
*/
package main

import (
	"fmt"
	"time"
)

func producer(nums ...int) <-chan int {
	channel := make(chan int)
	
	go func() {
		for _, num := range nums {
			channel <- num
		}
	}()
	return channel
}

func consumer(input <-chan int, quit <-chan bool) {
	for {
		select {
			case data := <- input:
				fmt.Printf("Data received: %d\n", data)
			case <- quit:
				fmt.Printf("Quit signalled\n")
				return
		}
	}
}

func main() {
	input := producer(1,2,3,4,5,)
	quit := make(chan bool) // This channel is used to signal the activity to quit

	go consumer(input, quit)
	time.Sleep(time.Duration(1 * time.Millisecond))
	
	// quit the consumer activity
	quit <- true

	time.Sleep(time.Duration(1 * time.Second))
}