package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessages(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessages() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	messages := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world",
	}

	for _, message := range messages {
		wg.Add(1)
		go updateMessages(message, &wg)
		wg.Wait()
		printMessages()
	}

	fmt.Println("Program stopped!")
}
