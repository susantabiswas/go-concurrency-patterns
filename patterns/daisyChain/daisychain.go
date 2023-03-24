/*
	Daisy Chain pattern is basically linking the channels together to form
	long stage pipelines.
*/
package main

import "fmt"

// Sends the content of right to left and adds 1 to it
// left <- (1 + right)
func link(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	// no. of channels to link
	n := 10000

	leftmost := make(chan int)

	left := leftmost
	right := leftmost

	// After the end of loop, we will have all the channels linked up
	// leftmost <- ....... <- right
	// then we feed 1 to the right most, which because of link function will
	// pass that to the left and also add 1 to it. 
	for i := 0; i < n; i++ {
		right = make(chan int)
		// left <- right channel link established
		// next time this right will be linked with a new right channel
		go link(left, right)
		left = right
	}
	
	right <- 1
	fmt.Printf("Leftmost channel received: %d\n", <-leftmost)
}