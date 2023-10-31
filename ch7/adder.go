package main

type Adder struct {
	start int
}

// methods are functions too
func (a Adder) AddTo(val int) int {
	return a.start + val
}
