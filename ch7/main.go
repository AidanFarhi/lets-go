package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func updateCounterWrong(c Counter) {
	// this is invoking Increment() on a copy of the original object
	c.Increment()
	fmt.Println("in updateCounterWrong:", c.String())
}

func updateCounterRight(c *Counter) {
	// this is invoking Increment() on a pointer to the original object
	c.Increment()
	fmt.Println("in updateCounterRight:", c.String())
}

// iotas
// rule of thumb - if the underlying value matters, don't use iota
// if you just want to differentiate between a set of values, use iota
type MailCategory int

// after the first line, every other line will increment
// the value of iota on each line
const (
	Uncategorized  MailCategory = iota // 0
	Personal                           // 1
	Spam                               // 2
	Social                             // 3
	Advertisements                     // 4
)

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
	c.Increment() // Go converts this to (&c).Increment() under the hood
	fmt.Println(c.String())

	// pointer receiver methods and functions
	var c2 Counter
	updateCounterWrong(c2)
	fmt.Println("in main:", c2.String())
	updateCounterRight(&c2)
	fmt.Println("in main:", c2.String())

	// using the tree code defined in int_tree.go
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(10)
	it = it.Insert(12)
	it = it.Insert(88)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))

	// methods can be treated like functions
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))

	f1 := myAdder.AddTo // this is called a method value
	fmt.Println(f1(5))

	f2 := Adder.AddTo // this is called a method expression
	fmt.Println(f2(myAdder, 5))

	// user defined types can use the operators of their underlying types
	type Person map[string]string
	p1 := Person{
		"name": "Bob",
		"age":  "20",
	}
	fmt.Println(p1["name"])
	fmt.Println(p1["age"])
	var p1Var map[string]string = p1
	var p1Var2 Person = p1
	fmt.Println(p1Var)
	fmt.Println(p1Var2)

	// using embedded types
	m := Manager{
		Employee: Employee{
			Name: "Bob",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	e1 := Employee{Name: "Jim", ID: "12346"}
	e2 := Employee{Name: "Joe", ID: "12347"}
	e3 := Employee{Name: "Bill", ID: "12348"}
	e4 := Employee{Name: "Frodo", ID: "99999"}
	m.Reports = append(m.Reports, e1)
	m.Reports = append(m.Reports, e2)
	m.Reports = append(m.Reports, e3)
	fmt.Println(m.Description())
	fmt.Println(m.HasReport(e1))
	fmt.Println(m.HasReport(e4))

	// using interfaces
	client := Client{
		// the SimpleStringProcessor conforms to the StringProcessor interface
		P: SimpleStringProcessor{},
	}
	client2 := Client{
		// the ComplexStringProcessor conforms to the StringProcessor interface
		P: ComplexStringProcessor{},
	}
	client.Program()
	client2.Program()

	// an empty interface indicates an optional value
	// or a variable that could store any type
	var myInterface interface{}
	myInterface = 20
	myInterface = "hello"
	myInterface = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}
	fmt.Println(myInterface)

	// reading data using an empty interface
	data := map[string]interface{}{}
	contents, err := os.ReadFile("test_data.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(contents, &data)
	fmt.Println(data)

	// type assertions
	type MyInt int
	var myInterface2 interface{}
	var myInt MyInt = 20
	myInterface2 = myInt
	myInt2 := myInterface2.(MyInt)
	fmt.Println(myInt2 + 10)
	myString, ok := myInterface2.(string)
	if !ok {
		fmt.Println(fmt.Errorf("unexpected type for %v", myInterface2))
	} else {
		fmt.Println(myString)
	}
}
