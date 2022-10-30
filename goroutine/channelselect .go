package goroutine

import (
	"fmt"
	"time"
)

// ############ channelとselect ############
func goroutine1(c chan string) {
	for {
		c <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutine2(c chan string) {
	for {
		c <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go goroutine1(c1)
	go goroutine2(c2)
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
