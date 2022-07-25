package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {

	// Set all this to 0
	// We need test the program fast
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	for indeks := 0; indeks < 100; indeks++ {
		main()
		if len(orderFinished) != 5 {
			t.Error("Wrong number of entries in orderFinished!")
		}
		orderFinished = []string{}
	}
}
