package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	p1 := MakePerson("bob", "johnson", 20)
	p2 := MakePersonPointer("bob", "johnson", 20)
	fmt.Println(p1, p2)
}
