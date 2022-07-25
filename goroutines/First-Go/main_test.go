package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printName2(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Write so the write files is not empty
	var wg sync.WaitGroup
	wg.Add(1)
	go printName2("Joshua Ryandafres Pangaribuan", &wg)
	wg.Wait()

	_ = w.Close()

	// Read the write files here
	result, _ := io.ReadAll(r)
	output := string(result)

	// Close the standard output
	os.Stdout = stdOut

	if !strings.Contains(output, "Joshua Ryandafres Pangaribuan") {
		t.Errorf("Expected JOshua Ryandafres Pangaribuan. But not there!")
	}
}
