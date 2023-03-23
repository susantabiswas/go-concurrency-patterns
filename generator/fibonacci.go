package main

import "fmt"

func fibonacciGenerator() chan int {
	fibChannel := make(chan int)
	
	go func() {
		for first, second := 0, 1; ;  {
			fibChannel <- first
			first, second = second, first + second 
		}
	}()

	return fibChannel
}

func FibonacciGeneratorDemo(n int) {
	fibChannel := fibonacciGenerator()

	// Get the first 5 fibonacci numbers
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-fibChannel)
	}
}