package pubadder_test

import (
	"pubadder"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 2
	actual := pubadder.Add(1, 1)
	if expected != actual {
		t.Errorf("expected: %d actual: %d", expected, actual)
	}
}

func TestSub(t *testing.T) {
	expected := 1
	actual := pubadder.Sub(2, 1)
	if expected != actual {
		t.Errorf("expected: %d actual: %d", expected, actual)
	}
}
