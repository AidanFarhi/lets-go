package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  *string
}

type Order struct {
	OrderId     int     `json:"orderId"`
	OrderAmount float64 `json:"orderAmount"`
	ProductId   int     `json:"productId"`
}

func makePointer[T any](t T) *T {
	return &t
}

func failedUpdate(g *int) {
	x := 10
	g = &x
	fmt.Println("g in failedUpdate:", g)
}

func failedUpdateV2(px *int) {
	x2 := 999
	px = &x2
}

func successfulUpdate(px *int) {
	*px = 100
}

// this is bad
func makePersonBad(p *Person) error {
	p.FirstName = "John"
	p.LastName = makePointer("Doe")
	return nil
}

// this is good
func makePersonGood() (Person, error) {
	p := Person{
		FirstName: "John",
		LastName:  makePointer("Doe"),
	}
	return p, nil
}

func main() {

	// pointer basics
	var x int32 = 10     // address: a1 | value: 10
	var y bool = true    // address: a2 | value: 1
	pointerX := &x       // address: a3 | value: a1
	pointerY := &y       // address: a4 | value: a2
	var pointerZ *string // address: a5 | value: nil
	fmt.Println(pointerX, pointerY, pointerZ)

	// the address operator
	s := "hello"
	pointerToS := &s
	fmt.Println(pointerToS)

	// the indirection operator
	sValue := *pointerToS
	fmt.Println(sValue)

	// go will panic if you try to dereference a nil pointer
	var nilPointer *int
	fmt.Println(nilPointer == nil)
	// fmt.Println(*nilPointer) -> panic

	// the pointer type
	num1 := 10
	var pointerToNum *int
	pointerToNum = &num1
	fmt.Println(*pointerToNum)

	// the new keyword
	num2 := new(int) // this returns a pointer to a zero-value instance of the type
	fmt.Println(num2 == nil)
	fmt.Println(*num2)

	// you cannot use & before primitive literals or a constant
	// because they exist only at compile time
	p1 := Person{
		FirstName: "Bob",
		// LastName:  "Bobson", this line will not compile
	}
	p2 := Person{
		FirstName: "Joe",
		LastName:  makePointer("Joeson"),
	}
	fmt.Println(p1)
	fmt.Println(p2)

	// when pointers are passed to functions, the function gets
	// a copy of the pointer. This means the original data can
	// be modified by the called function.
	var f *int
	failedUpdate(f)
	fmt.Println(f)

	xVal := 99
	failedUpdateV2(&xVal)
	fmt.Println(xVal)
	successfulUpdate(&xVal)
	fmt.Println(xVal)

	// the only time you should use pointer parameters to modify a variable
	// is when the function expects an interface.
	order := Order{}
	data := []byte(`{"orderId": 1234, "orderAmount": 23.44, "productId": 5542}`)
	json.Unmarshal(data, &order)
	fmt.Println(order)
}
