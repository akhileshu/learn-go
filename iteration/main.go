package main

import "strings"

func main() {
	// entry point
}
const repeatCount = 5
func Repeat(s string)  string {
	// Strings in Go are immutable , so we have to create a new string
	// using the strings.Builder to minimize memory copying
	var result strings.Builder
	for range repeatCount {
		result .WriteString(s)
	}
	return result.String()
}
