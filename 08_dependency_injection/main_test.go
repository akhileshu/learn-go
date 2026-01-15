package main

import (
	"bytes"
	"go-playground/package/testing_utils"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	testing_utils.AssertStrings(t, got, want)
}
