package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	// Run main function here
	main()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "Rp.2184000,00") {
		t.Error("Balance not equal! probably race conditions!")
	}
}
