package Integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected %d, but got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(7, 2)
	fmt.Println(sum)
	// Output: 9
}
