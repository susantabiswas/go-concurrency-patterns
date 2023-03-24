package main
import "fmt"

func main() {
	fmt.Println("** Fibonacci numbers using generator pattern **")
	FibonacciGeneratorDemo(6)

	fmt.Println("\n** Integer numbers using generator pattern **")
	NumberGeneratorDemo()
}