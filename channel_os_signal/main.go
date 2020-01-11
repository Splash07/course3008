package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Print: hello 1, hello 2, hello 3... in every 800ms in go routine
	// The key here is to keep the MAIN routine alive (active/sleep -> not finished). All goroutines are terminated
	//		as soon as the MAIN routine is terminated
	forever := make(chan os.Signal)
	signal.Notify(forever, os.Interrupt)
	go func() {
		i := 0
		for true {
			i++
			fmt.Println("Hello ", i)
			time.Sleep(time.Millisecond * 800)
		}
	}()

	<-forever //routine -> sleep/pending
	// then, notify: receive from channel -> routine active again
	fmt.Println("Program end. Thanks for using this program")
}
