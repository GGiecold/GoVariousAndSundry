package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Not within a goroutine.")
	go func() {
		fmt.Println("Within a goroutine.")
	}()
	fmt.Println("Not within a goroutine.")
	runtime.Gosched()
}
