package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			chan1 <- "ping"
		}
	}()
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			chan2 <- "pong"
		}
	}()

	for i := 0; i < 6; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}
}
