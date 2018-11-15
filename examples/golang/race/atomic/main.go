package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var count int64

	for i := 0; i < 50; i++ {
		go func(no int) {
			for i := 0; i < 500; i++ {
				atomic.AddInt64(&count, 1)
			}
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(atomic.LoadInt64(&count))
}
