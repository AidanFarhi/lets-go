package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changePersonName(person Person) {
	person.name = "jim"
}

func changePersonNameV2(person *Person) {
	person.name = "jim"
}

func changePerson(person *Person) {
	*person = Person{"frodo", 999}
}

func updateInt(intPointer *int) {
	*intPointer = 99
}

func addToSlice(slice []int) {
	numsToAdd := []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, numsToAdd...)
	fmt.Println(slice)
}

func main() {
	/*
		- Pointers -

		Use them sparingly. Be nice to your garbage collector.
	*/
	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string
	fmt.Println(x)
	fmt.Println(pointerX)
	fmt.Println(y)
	fmt.Println(pointerY)
	fmt.Println(pointerZ)

	// address operator
	s := "hello"
	sPointer := &s
	fmt.Println(s, sPointer)

	// indirection operator (derefrencing)
	sPointerValue := *sPointer
	fmt.Println(sPointerValue)

	// passing pointers to functions
	p1 := Person{"bob", 20}
	fmt.Println(p1.name)

	changePersonName(p1)
	fmt.Println(p1.name)

	changePersonNameV2(&p1)
	fmt.Println(p1.name)

	intVar := 10
	updateInt(&intVar)
	fmt.Println(intVar)

	fmt.Println(p1)
	changePerson(&p1)
	fmt.Println(p1)

	// fun with slices
	originalSlice := []int{11, 22, 33}
	addToSlice(originalSlice)
	fmt.Println(originalSlice)
}
