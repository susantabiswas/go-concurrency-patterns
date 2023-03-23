/*
	FAN IN
	Use case: When you want to combine multiple item lists together in a single item list

	FAN-OUT
	Use case: When you want to branch out multiple item lists from a single item list
*/
package main

import (
	"fmt"
	"sync"
)

// Creates a read-only channel from an array
func extract(nums []int) <- chan int {
	channel := make(chan int)

	go func() {
		for _, num := range nums {
			channel <- num
		}
		close(channel)
	}()

	return channel
}

// Transform operation on the data
func transform(input <-chan int, workType int) <- chan string {
	transformed := make(chan string)

	go func() {
		for num := range input {
			transformed <- fmt.Sprintf("Performed work %d on number: %d", workType, num)
		}
		close(transformed)
	}()

	return transformed
}

// Loads the data from different channels to a common channel
func load(inputs ...<-chan string) <-chan string {
	merged := make(chan string)
	
	var wg sync.WaitGroup
	wg.Add(len(inputs))

	// Loads the data from a channel to the common channel
	loadData := func(ch <-chan string) {
		for data := range ch {
			merged <- data
		}
		wg.Done()
	}

	// Start the data loading operation from the
	// different channels
	for _, channel := range inputs {
		go loadData(channel)
	}

	// To close the channel, ensure that all the goroutines are finished
	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func useCaseOneDemo() {
	// Initial data source
	data := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// create a channel from the initial source
	input := extract(data)

	// FAN-OUT
	// From the single source, create fan-out channels
	// Data sources: Fan-out channels.
	// Let's say we want to fan-out to perform different operations on them
	ch1 := transform(input, 1)
	ch2 := transform(input, 2)

	// FAN-IN
	// Combine the transformed data from the different sources
	outputChannel := load(ch1, ch2)

	// Consume the data.
	for data := range outputChannel {
		fmt.Printf("%s\n", data)
	}
}

func main() {
	useCaseOneDemo()
}