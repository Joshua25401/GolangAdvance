package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

// Add mutex parameter!
func updateMessage(s string, mtx *sync.Mutex) {
	defer wg.Done()

	// Lock this potential race conditions
	// -> by using mtx.Lock()
	mtx.Lock()
	msg = s
	// Don't forget to Unlock()
	// -> by using mtx.Unlock()
	mtx.Unlock()
}

func main() {
	msg = "Hello, Joshua!"

	// Declare mutex here
	// -> by writing sync.Mutex
	var mtx sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, Chesya!", &mtx)
	go updateMessage("Hello, Somebody!", &mtx)
	wg.Wait()

	// Print result
	fmt.Println(msg)
	fmt.Println("Program stopped!")
}
