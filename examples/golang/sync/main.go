package main

import (
	"fmt"
	"sync"
)

var s []int

func add(n int) {
	s = append(s, n)
}

func main() {
	o := sync.Once{}
	exit := make(chan int, 2)

	go func() {
		for i := 0; i < 50; i++ {
			add(0)
		}
		exit <- 1
	}()

	o.Do(func() {
		for i := 0; i < 50; i++ {
			add(1)
		}
		exit <- 1
	})
	<-exit
	<-exit
	fmt.Println(s)
}
