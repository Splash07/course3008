package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go Task1(&wg)
	go Task2(&wg)
	go Task3(&wg)
	wg.Wait()
}

func Task1(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Millisecond)
	fmt.Println("task1 done")
	defer wg.Done()
}

func Task2(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Millisecond)
	fmt.Println("task2 done")
	defer wg.Done()
}

func Task3(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Millisecond)
	fmt.Println("task3 done")
	defer wg.Done()
}
