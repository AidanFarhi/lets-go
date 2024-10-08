package main

import "testing"

func TestAddNumbers(t *testing.T) {
	result := AddNumbers(2, 3)
	if result != 5 {
		t.Error("expected: 5 got:", result)
	}
}
