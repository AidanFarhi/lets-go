package main

import (
	"ch8/stack"
	"ch8/util"
	"fmt"
)

func main() {

	s := stack.Stack[string]{}
	s.Push("first")
	s.Push("second")
	s.Push("third")
	fmt.Println(s.Contains("second"))
	for s.Size() > 0 {
		fmt.Println(s.Pop())
	}
	fmt.Println(s.Contains("second"))

	words := []string{"one", "two", "three", "four"}
	filtered := util.Filter(words, func(s string) bool {
		return s != "three"
	})
	fmt.Println(filtered)
	lengths := util.Map(filtered, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)
	sum := util.Reduce(lengths, 0, func(acc, val int) int {
		return acc + val
	})
	fmt.Println(sum)
}
