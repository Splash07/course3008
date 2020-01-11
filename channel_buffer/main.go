package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffered channel with capacity of 2

	// len <= capacity -> routine doesn't sleep 
	ch <- 1
	ch <- 2
	ch <- 3

	// len > capacity -> routine sleep
	fmt.Println("Read channel")
}
