package main

import (
	"fmt"
)

func main() {
	memory := &Memory{
		Channel: make(chan string, 1),
	}

	go memory.ReadChannel()
	go memory.MagicChannel()

	msg := ""

	for {
		msg = ""
		fmt.Scanln(&msg)

		if msg == "" {
			break
		}

		go memory.SendChannel(msg)
	}
}
