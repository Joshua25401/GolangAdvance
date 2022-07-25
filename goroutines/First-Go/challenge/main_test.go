package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessages(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessages("Josh", &wg)
	wg.Wait()

	if msg != "Josh" {
		fmt.Println("Expected Josh but not found!")
	}
}

func Test_printMessages(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Write messages here
	msg = "Hello, World!"
	printMessages()

	_ = w.Close()
	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, msg) {
		t.Errorf("Expected %v. but not found!", msg)
	}
}
