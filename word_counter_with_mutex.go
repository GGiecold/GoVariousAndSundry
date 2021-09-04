package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type words struct {
	found map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, tally int) {
	count, ok := w.found[word]
	if !ok {
		w.found[word] = tally
		return
	}
	w.found[word] = count + tally
}

// Open a file specified by |filename|, scan its content,
// tallying the words therein and updating a |dict|.
func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}

func main() {
	var wait_group sync.WaitGroup
	words := newWords()
	for _, filename := range os.Args[1:] {
		wait_group.Add(1)
		go func(file string) {
			if err := tallyWords(file, words); err != nil {
				fmt.Println(err.Error())
			}
			wait_group.Done()
		}(filename)
	}
	wait_group.Wait()
	fmt.Println("Here are the words appearing more than once:\n")
	for word, count := range words.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}
