package main

import (
	"fmt"
	"sync"
)

/*
sync.Mutex
	- Mutex stands for 'Mutual Exclusion' that allows us to deal with race conditions
	- Simple to use
	- Dealing with shared resources and concurrent/parallel goroutines
	- Lock/Un-Lock
	- Test for race conditions when running code, or testing it

Race Conditions
	- The conditions when multiple goroutines try to access the same data
	- Can be difficult to spot when reading code
	- Go can check this condition when running a program or when testing the code.
*/

// Shared Data between goroutines
// Race conditions happen to this shared data
var msg string

// WaitGroup
var wg sync.WaitGroup

// Function to update the shared data
func updateMessage(s string) {
	defer wg.Done()
	// Update shared data here
	// It becomes race conditions when multiple goroutines
	// at the same time change the msg variable or the shared variable
	msg = s
}

func main() {
	msg = "Hello, Joshua!"

	wg.Add(2)
	// Spawn goroutines
	go updateMessage("Hello, Chesya!")
	// Spawn another goroutines
	go updateMessage("Hello, Somebody!")
	wg.Wait()

	// Run it, and see the RACE CONDITIONS!!
	// we can execute this by using :
	//	- make test_race command (if you want to test race conditions in go)
	// 	- make run_normal command
	fmt.Println(msg)
}
