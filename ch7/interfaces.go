package main

import "fmt"

type Logic interface {
	Process(data string) string
}

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return fmt.Sprintf("This is your data: %s", data)
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	data := "{name: bob, age: 44}"
	result := c.L.Process(data)
	fmt.Println(result)
}
