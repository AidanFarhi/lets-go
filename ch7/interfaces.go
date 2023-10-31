package main

import "fmt"

// the SimpleStringProcessor meets the StringProcessor
// interface implicitly.
type SimpleStringProcessor struct{}

func (ssp SimpleStringProcessor) Process(data string) string {
	fmt.Println("processing:", data)
	return "done"
}

// the ComplexStringProcessor meets the StringProcessor
// interface implicitly.
type ComplexStringProcessor struct{}

func (csp ComplexStringProcessor) Process(data string) string {
	fmt.Println("complex processing:", data)
	return "done"
}

type StringProcessor interface {
	Process(data string) string
}

type Client struct {
	P StringProcessor
}

func (c Client) Program() {
	data := "this is the data"
	c.P.Process(data)
}
