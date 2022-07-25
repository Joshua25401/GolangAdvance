package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%v\n", strings.ToUpper(s))
	}
}

/*
	Channel :
	- Is the way we can push data to goroutine
	- Or getting data from other goroutine

	Syntax
	- chan interface{}
*/

func main() {
	ping := make(chan string)
	pong := make(chan string)

	// Use defer keyword to close the channel
	defer close(ping)
	defer close(pong)

	// This shout() function will running in the background
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")
	for {
		// Print a prompt
		fmt.Print("->")
		// Get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		// Sending data to ping channel
		ping <- userInput

		// Wait for response
		response := <-pong
		fmt.Println("You typed :", response)
	}

	fmt.Println("All done closing channel!")
}
