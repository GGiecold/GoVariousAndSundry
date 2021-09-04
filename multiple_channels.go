package main

import (
	"fmt"
	"os"
	"time"
)

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		// Copy input stream to data.
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}

func main() {
	timed_out := time.After(30 * time.Second)
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-timed_out:
			fmt.Println("Time out!")
			os.Exit(0)
		}
	}
}
