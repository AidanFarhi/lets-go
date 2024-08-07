package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Incrementer interface {
	Increment()
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type Logic interface {
	Process(data string) string
}

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return data + " -> processed"
}

type Client struct {
	Logic
}

func (c Client) Program(data string) {
	result := c.Process(data)
	fmt.Println(result)
}

// interfaces are comparable
type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func main() {

	var myStringer fmt.Stringer
	var myIncrementer Incrementer

	// method set:
	// Increment
	// String
	pointerCounter := &Counter{}
	// method set:
	// String
	valueCounter := Counter{}

	myStringer = pointerCounter
	myStringer = valueCounter
	myIncrementer = pointerCounter
	// this will not compile because the Increment method
	// is only part of the method set of pointer instances.
	// myIncrementer = valueCounter

	fmt.Println(myStringer)
	fmt.Println(myIncrementer)

	c := Client{
		Logic: LogicProvider{},
	}
	c.Program("data")

	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	// var dis2 = DoubleIntSlice{1, 2, 3}

	// the types match but the pointers do not
	DoublerCompare(&di, &di2)
	// the types do not match
	DoublerCompare(&di, dis)
	// this will panic at runtime because
	// slices are not comparable
	// DoublerCompare(dis, dis2)

	// the empty interface says nothing
	// an empty interface states that the
	// variable can store any value whose
	// type implements zero or more methods.
	// this matches every type in go.
	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		Name string
		Age  int
	}{"Bob", 20}
	fmt.Println(i)

	// any is a type alias for interface{}
	data := map[string]any{}
	contents, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("error reading file")
	}
	json.Unmarshal(contents, &data)
	fmt.Println(data)
}
