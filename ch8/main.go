package main

import (
	"ch8/stack"
	"fmt"
)

func main() {
	s := stack.Stack[string]{}
	s.Push("first")
	s.Push("second")
	s.Push("third")
	for s.Size() > 0 {
		fmt.Println(s.Pop())
	}
}
