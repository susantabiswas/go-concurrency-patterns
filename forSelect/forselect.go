/*
	Select can be used to run cases for different channels. Now the same
	select construct can be used inside a for loop. Using the combination of these
	two, we can have like an infinite loop listening for receiving events.
*/
package main

import "fmt"

func producer(nums ...int) <-chan int {
	channel := make(chan int)
	
	go func() {
		for _, num := range nums {
			channel <- num
		}
		close(channel)
	}()

	return channel
}

func consumer(input <-chan int) {
	// variant with finite loop
	for i := 0; i < 2; i++ {
		fmt.Printf("Data received from finite loop: %d\n", <- input)
	}

	// variant with infinite loop
	for {
		select {
			case data, ok := <- input:
				if !ok {
					return
				}
				fmt.Printf("Data received from infinite loop: %d\n", data)
			default:
		}
	}
}

func main() {
	producerChannel := producer(1,2,3,4,5,)

	consumer(producerChannel)
}