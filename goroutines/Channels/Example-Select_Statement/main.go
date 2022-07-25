package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		// Sleep for 6 seconds
		time.Sleep(6 * time.Second)
		ch <- "This is from server1 :)"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server2 :("
	}
}

func main() {
	fmt.Println("SELECT with channels")
	fmt.Println("====================")

	// Create channel
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Don't forget to close the channel
	defer close(channel1)
	defer close(channel2)

	// Spawn the goroutine to running both sever
	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Println("Case 1 returning :", s1)
		case s2 := <-channel1:
			fmt.Println("Case 2 returning :", s2)

		case s3 := <-channel2:
			fmt.Println("Case 3 returning :", s3)

		case s4 := <-channel2:
			fmt.Println("Case 4 returning :", s4)

		default:
			// Avoiding deadlock
		}
	}
}
