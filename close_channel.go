package main

import (
	"fmt"
	"time"
)

// |ch| is a received channel.
// |done| is a sending channel.
func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("I say goodbye!")
			close(ch)
			return
		default:
			ch <- "You say hello!"
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	msg := make(chan string)
	done := make(chan bool)
	until := time.After(5 * time.Second)
	go send(msg, done)
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}
