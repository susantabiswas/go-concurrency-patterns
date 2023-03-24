/*
	Range-Close

	This pattern is variant of classic producer-consumer pattern.
	Here the producer sends the data to the channel and also closes
	the channel to indicate that it no longer wants to send more.

	The consumer can keep receiving data from producer channel until
	producer has signalled that it will no longer publish data.
*/
package main

import "fmt"

// Producer: Keeps generating data 
func producer(n int) chan int {
	channel := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			channel <- i
		}
		// close the channel to indicate end of data production
		close(channel)
	}()

	return channel
}

// Consumer: Keeps consuming data from the producer till it gets
func consumer(producer <-chan int) {
	for num := range producer {
		fmt.Printf("Consumer %d from producer\n", num)
	}
	fmt.Printf("Consumer stopped receiving from producer\n")
}

func main() {
	producerChannel := producer(5)
	consumer(producerChannel)
}