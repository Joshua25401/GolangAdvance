package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "Hello"

	// Example of builder
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")

	// Reuse the builder
	sb.Reset()
	words := []string{"Hello", "World"}
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")

	fmt.Println(sb.String())
}
