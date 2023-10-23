package main

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}
