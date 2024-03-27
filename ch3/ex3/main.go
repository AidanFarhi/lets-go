package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {

	emp1 := Employee{"Bob", "Smith", 123}

	emp2 := Employee{
		firstName: "Jim",
		lastName:  "Johnson",
		id:        456,
	}

	var emp3 = Employee{}
	emp3.firstName = "Bill"
	emp3.lastName = "Smith"
	emp3.id = 789

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
