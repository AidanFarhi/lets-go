package adder

import "testing"

func Test_adder(t *testing.T) {
	result := AddNums(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
