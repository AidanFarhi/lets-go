package main

import "fmt"

func updateCounterWrong(c Counter) {
	c.Increment()
	fmt.Println("in updateCounterWrong:", c.String())
}

func updateCounterRight(c *Counter) {
	c.Increment()
	fmt.Println("in updateCounterRight:", c.String())
}

func main() {
	// invoke a method on a struct
	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       65,
	}
	fmt.Println(p.String())

	// using pointer receiver methods
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	// pointer receiver methods and functions
	var c2 Counter
	updateCounterWrong(c2)
	fmt.Println("in main:", c2.String())
	updateCounterRight(&c2)
	fmt.Println("in main:", c2.String())
}
