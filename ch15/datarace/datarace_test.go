package datarace

import "testing"

func TestGetCounter(t *testing.T) {
	expected := 5000
	actual := GetCounter()
	if expected != actual {
		t.Errorf("expected: %d | actual: %d", expected, actual)
	}
}
