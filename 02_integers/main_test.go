package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := AddInt(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

/* func ExampleAddInt() {
	sum := AddInt(1, 5)
	fmt.Println(sum)
	// Output: 6
}
*/

func ExampleAddInt() {
	fmt.Println(AddInt(1, 5))
	// Output: 6
}
