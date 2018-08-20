package main

import (
	"fmt"
	"time"
)

func Producer(id int, item chan int) {
	for i:=0; i<10; i++ {
		item <- i
		fmt.Printf("Producer %d produces data: %d\n", id, i)
		time.Sleep(10 * 1e6)
	}
}

func Consumer(id int, item chan int) {
	// for i:=0; i<20; i++ {
	c_item := <-item
	fmt.Printf("Consumer %d get data: %d\n", id, c_item)
	time.Sleep(10 * 1e6)
	// }
}

func main() {
	item := make(chan int, 6)
	go Producer(1, item)
	go Producer(2, item)
	for i:=0; i<20; i++ {
		go Consumer(1, item)
	}
	time.Sleep(1 * 1e9)
}
