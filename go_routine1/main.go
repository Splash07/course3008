package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go Task1(c)
	go Task2(c)
	go Task3(c)
	time.Sleep(2 * time.Millisecond)
}

func Task1(c chan int) {
	fmt.Println("task1 done")
	c <- 1
}

func Task2(c chan int) {
	<-c
	fmt.Println("task2 done")
	c <- 2
}

func Task3(c chan int) {
	<-c
	fmt.Println("task3 done")
}
