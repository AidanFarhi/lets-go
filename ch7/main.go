package main

import (
	"fmt"
)

func main() {
	p1 := Person{
		FirstName: "bob",
		LastName:  "smith",
		Age:       34,
	}
	fmt.Println(p1.String())

	myCounter := Counter{}
	myCounter.Increment()
	myCounter.Increment()
	fmt.Println(myCounter.String())

	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false

	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))

	adderMethod := myAdder.AddTo
	fmt.Println(adderMethod(10))

	myOtherAdder := Adder{start: 100}
	adderMethodFromType := Adder.AddTo
	fmt.Println(adderMethodFromType(myOtherAdder, 10))

	fmt.Println(Adder.AddTo(myOtherAdder, 100))

	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Description())

	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.Double())
	fmt.Println(o.IntPrinter(10))

	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
