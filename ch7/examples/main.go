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

// methods are functions too
type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

// you can declare types based on other types
type HighScore Score

// iotas
type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

// use embedding for composition
type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

// all fields or methods declared on an
// embedded field are promoted to the containing struct
type Manager struct {
	Employee // this is an embedded field
	Reports  []Employee
}

func (m Manager) DisplayReports() {
	for _, e := range m.Reports {
		fmt.Println(e.Description())
	}
}

// if a containing struct has fields or methods with
// the same name as the embedded field, you need to
// use the embedded field's type to refer to the
// obscured fields or methods.
type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

// there is no dynamic dispatch
type InnerA struct {
	A int
}

func (i InnerA) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i InnerA) Double() string {
	return i.IntPrinter(i.A * 2)
}

type OuterA struct {
	InnerA
	S string
}

func (o OuterA) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
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

	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))
	// method value
	f1 := myAdder.AddTo
	fmt.Println(f1(5))
	// method expression
	f2 := Adder.AddTo
	fmt.Println(f2(Adder{start: 100}, 15))

	// assigning untyped constants is valid
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	// hs = s
	// s = i
	s = Score(i)
	hs = HighScore(s)
	fmt.Println(hs)
	var newScore Score = 50
	scoreWithBonus := newScore + 100 // type of scoreWithBonus is Score
	fmt.Println(scoreWithBonus)

	manager := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(manager.ID)
	fmt.Println(manager.Description())
	manager.Reports = append(manager.Reports, Employee{"Jim Jimson", "44432"})
	manager.Reports = append(manager.Reports, Employee{"Steve Jobs", "85854"})
	manager.DisplayReports()

	o := Outer{
		Inner: Inner{
			X: 100,
		},
		X: 200,
	}
	fmt.Println(o.X)
	fmt.Println(o.Inner.X)

	o2 := OuterA{
		InnerA: InnerA{
			A: 10,
		},
		S: "hello",
	}
	fmt.Println(o2.Double())
}
