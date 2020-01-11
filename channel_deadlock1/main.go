package main

func main() {
	ch := make(chan int)

	// read channel cause current goroutine to sleep
	<- ch
}
