package main

import (
	"fmt"
	"time"
)

func add(ch chan int) {
	for i := 0; i < 50; i++ {
		ch <- <-ch + 1
	}
}

func main() {
	ch := make(chan int, 3)
	ch <- 0

	for i := 0; i < 500; i++ {
		go add(ch)
	}

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
