package main

import "fmt"

func main() {
	ch := make(chan int) // unbuffered channel

	// write to unbuffered channel also causes routine to sleep

	go func() {
		ch <- 2
	}()

	fmt.Println("program end")
}


// Two way to fix this deadlock situation
// 1st way is to have another gorountine write to unbuffered channel 
// 2nd way is to use buffered channel