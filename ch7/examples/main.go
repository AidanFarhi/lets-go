package main

import (
	"fmt"
	"time"
)

// user defined types
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// a method for Person
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// types can be primitive or compound literal types
type Score int
type ConvertScore func(string) Score
type TeamScores map[string]Score

type Counter struct {
	total       int
	lastUpdated time.Time
}

// use a pointer receiver if the method mutates the object
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c)
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c)
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {

	p := Person{
		FirstName: "bob",
		LastName:  "bobson",
		Age:       43,
	}
	fmt.Println(p)

	var c Counter
	fmt.Println(c)
	c.Increment()
	fmt.Println(c)

	c2 := &Counter{}
	fmt.Println(c2)
	c2.Increment()
	fmt.Println(c2)

	var c3 Counter
	doUpdateWrong(c3)
	fmt.Println("in main:", c3)
	doUpdateRight(&c3)
	fmt.Println("in main:", c3)
}
