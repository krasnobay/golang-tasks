package main

import (
	"fmt"
	"golang-tasks/examples/utils"
	"time"
)

//Memory data store
type Memory struct {
	Channel chan string
}

//ReadChannel gorutines reading data from the channel
func (m Memory) ReadChannel() {
	for {
		msg := <-m.Channel
		fmt.Println("Read: " + msg)
	}
}

//SendChannel send msg to Channel
func (m *Memory) SendChannel(msg string) {
	m.Channel <- msg
}

//MagicChannel magic traffic generation
func (m *Memory) MagicChannel() {
	for {
		m.Channel <- m.generateMagic()
		time.Sleep(time.Second * 1)
	}
}

func (m Memory) generateMagic() string {
	return utils.GetMD5Hash(time.Now().String())
}
