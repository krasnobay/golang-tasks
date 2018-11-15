package main

import (
	"fmt"
	"sync"
	"time"
)

func add(count *int, mu *sync.Mutex) {
	for i := 0; i < 50; i++ {
		mu.Lock()
		*count += 1
		mu.Unlock()
	}
}

func main() {
	var mu sync.Mutex
	count := 0

	for i := 0; i < 500; i++ {
		go add(&count, &mu)
	}

	time.Sleep(time.Second)

	mu.Lock()
	fmt.Println(count)
	mu.Unlock()
}
