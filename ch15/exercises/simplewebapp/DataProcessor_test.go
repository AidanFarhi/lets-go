package main

import "testing"

func Test_parser(t *testing.T) {
	data := []struct {
		id     string
		op     string
		val1   int
		val2   int
		errMsg string
	}{
		{"1", "+", 2, 3, nil},
	}
}
