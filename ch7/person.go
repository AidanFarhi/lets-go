package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// String
// This a method that returns a user-friendly string representation
// of an object of type Person.
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
