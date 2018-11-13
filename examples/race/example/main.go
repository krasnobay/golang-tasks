package main

import (
	"fmt"
)

func main() {
	count := 0

	for i := 0; i < 50; i++ {
		go func(no int) {
			for i := 0; i < 500; i++ {
				count += 1
			}
		}(i)
	}

	fmt.Println(count)
}
