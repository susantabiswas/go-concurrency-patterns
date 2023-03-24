/*
	Sequence pattern allows to coordinate the timing/order of execution of work.

	Lets say we want to ensure that different goroutines execute something and then
	only perform the next batch of job when all are done with the work.
	Eg: Copying a large file in chunks, we allow the chunks to be copied by go routines and
	make all the goroutines wait once they have copied so that we commit the changes.
*/
package main

import (
	"fmt"
)

// Represents work
type Work struct {
	workType string // Type of work eg electrical, appliance setup
	readyForNext chan bool // Whether is ready to start next work
}

// Represents work done by a worker
func doWork(work string) <-chan Work {
	worker := make(chan Work) // channel for work done
	ready := make(chan bool) // channel to indicate whether to wait or not

	// This will keep performing work of a particular type
	go func() {
		for i := 0; ; i++ {
			worker <- Work{work, ready}
			<-ready // Wait till the other worker has also finished the work
		}
	}()

	return worker
}

// Merge the work details from both the workers in a single
// channel to avoid any blocking calls in the main
// The common channel will accept work details from both the 
// workers
func fanIn(worker1, worker2 <-chan Work) <- chan Work {
	channel := make(chan Work)

	go func() {
		for {
			channel <- <-worker1
		}
	}()

	go func() {
		for {
			channel <- <-worker2
		}
	}()

	return channel
}

func main() {
	// Worker1 does "electrical" and worker2 does "appliance setup" work 
	workChannel := fanIn(doWork("electrical"), doWork("appliance setup"))

	// Both will only start working on the next building when
	// both the electrical wiring work and appliances are setup for the
	// current building
	for building := 1; building <= 3; building++ {
		work1 := <-workChannel
		work2 := <-workChannel
		fmt.Printf("Work %s done for building %d\n", work1.workType, building)
		fmt.Printf("Work %s done for building %d\n", work2.workType, building)
		
		// signal readiness for the work of next building
		work1.readyForNext <- true
		work2.readyForNext <- true
	}
}