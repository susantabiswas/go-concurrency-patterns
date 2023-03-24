package main

import (
	"fmt"
)

/* Generates numbers infinitely

	A go routine keeps generating numbers infinitely and
	pushes them to the channel. Since the channel push 
	is a blocking call, the code only executes when the number is consumed. 
*/
func numGenerator() chan int {
	numChannel := make(chan int)
	
	go func() {
		for i := 0; ; i++ {
			numChannel <- i
		}
	}()

	return numChannel
}

func NumberGeneratorDemo() {
	numChannel := numGenerator()

	// Consume the numbers whenever needed
	fmt.Printf("Getting a number for the first time: %d\n", <-numChannel)
	fmt.Printf("Getting a number for the second time: %d\n", <-numChannel)

	for i := 1; i <= 3; i++ {
		fmt.Printf("Getting a number using loop: %d\n", <-numChannel)
	}
}

