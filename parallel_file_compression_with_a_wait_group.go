package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()
	gz_writer := gzip.NewWriter(out)
	_, err = io.Copy(gz_writer, in)
	gz_writer.Close()
	return err
}

func main() {
	var wait_group sync.WaitGroup
	var file_cnt int = -1
	var file string
	for file_cnt, file = range os.Args[1:] {
		wait_group.Add(1)
		go func(filename string) {
			compress(filename)
			wait_group.Done()
		}(file)
	}
	wait_group.Wait()
	fmt.Printf("Compressed %d files\n", file_cnt+1)
}
