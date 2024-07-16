package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {

	emp1 := Employee{
		"joe",
		"smith",
		123,
	}

	emp2 := Employee{
		firstName: "jim",
		lastName:  "johnson",
		id:        345,
	}

	var emp3 Employee
	emp3.firstName = "bob"
	emp3.lastName = "jones"
	emp3.id = 654

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
