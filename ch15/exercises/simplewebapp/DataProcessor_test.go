package main

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_parser(t *testing.T) {
	data := []struct {
		id         string
		expression []byte
		expected   Input
		err        error
	}{
		{"add", []byte("1\n+\n2\n3"), Input{"1", "+", 2, 3}, nil},
		{"subtract", []byte("2\n-\n2\n3"), Input{"2", "-", 2, 3}, nil},
		{"multiply", []byte("3\n*\n2\n3"), Input{"3", "*", 2, 3}, nil},
		{"divide", []byte("4\n/\n10\n2"), Input{"4", "/", 10, 2}, nil},
	}
	for _, d := range data {
		expected := d.expected
		t.Run(d.id, func(t *testing.T) {
			actual, _ := parser(d.expression)
			if diff := cmp.Diff(expected, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestDataProcessor(t *testing.T) {
	inputData := []struct {
		id    string
		input []byte
	}{
		{"1", []byte("1\n+\n2\n3")},
		{"2", []byte("2\n-\n2\n3")},
		{"3", []byte("3\n*\n2\n3")},
		{"4", []byte("4\n/\n10\n2")},
	}
	expectedResults := map[string]Result{
		"1": Result{"1", 5},
		"2": Result{"2", -1},
		"3": Result{"3", 6},
		"4": Result{"4", 5},
	}
	in := make(chan []byte, 4)
	out := make(chan Result, 4)
	for _, d := range inputData {
		in <- d.input
	}
	close(in)
	DataProcessor(in, out)
	for r := range out {
		expected := expectedResults[r.Id]
		t.Run(r.Id, func(t *testing.T) {
			actual := r
			if diff := cmp.Diff(expected, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestWriteData(t *testing.T) {
	in := make(chan Result, 1)
	in <- Result{"1", 100}
	close(in)
	buf := bytes.NewBuffer([]byte{})
	WriteData(in, buf)
	expected := "1:100\n"
	actual := buf.String()
	if expected != actual {
		t.Errorf("expected: %s | actual: %s\n", expected, actual)
	}
}
