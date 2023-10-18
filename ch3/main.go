package main

import "fmt"

func main() {
	/*
		- Arrays -

		Arrays are very rigid in Go. When you declare an array in go,
		it's size is part of it's type. There is no way to convert
		an array of one type to another.

		This means you cannot create functions that work with arrays of
		any size and you cannot assign arrays of different sizes to the same
		variable.
	*/
	fmt.Println("----------- Arrays -----------")
	var intArray [3]int
	var intArrayLiteral = [3]int{1, 2, 3}
	var intSparseArray = [10]int{100, 3: 200, 300}
	var intArrayLiteralAlternate = [...]int{1, 2, 3}
	var matrix = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(intArray)
	fmt.Println(intArrayLiteral)
	fmt.Println(intSparseArray)
	fmt.Println(intArrayLiteralAlternate)
	fmt.Println(intArrayLiteral == intArrayLiteralAlternate)
	fmt.Println(matrix)
	fmt.Println(matrix[0][1])
	fmt.Println(len(intSparseArray))

	/*
		- Slices -

		A data structure that holds a sequence of values.

		  - Capacity

		  Every slice has a capacity, which is the number of
		  consecutive memory locations reserved. Once the capacity
		  is met, the Go runtime will allocate a new slice with a
		  larger capacity.
	*/
	fmt.Println("----------- Slices -----------")
	var mySlice = []int{10, 20, 30}
	var myOtherSlice = []int{11, 22, 33}
	var nilSlice []int
	mySlice = append(mySlice, 40, 50)
	myOtherSlice = append(myOtherSlice, mySlice...)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2])
	fmt.Println(nilSlice)
	fmt.Println(nilSlice == nil)
	fmt.Println(len(mySlice))
	fmt.Println(mySlice)
	fmt.Println(myOtherSlice)

	// Capacity
	capSlice := []int{}
	fmt.Println(capSlice, len(capSlice), cap(capSlice))
	capSlice = append(capSlice, 10)
	fmt.Println(capSlice, len(capSlice), cap(capSlice))
	capSlice = append(capSlice, 20)
	fmt.Println(capSlice, len(capSlice), cap(capSlice))
	capSlice = append(capSlice, 30)
	fmt.Println(capSlice, len(capSlice), cap(capSlice))
	capSlice = append(capSlice, 40)
	fmt.Println(capSlice, len(capSlice), cap(capSlice))
	capSlice = append(capSlice, 50)
	fmt.Println(capSlice, len(capSlice), cap(capSlice))
	capSlice = append(capSlice, 60)
	fmt.Println(capSlice, len(capSlice), cap(capSlice))

	// make
	makeSlice := make([]int, 5) // specifying capacity
	makeSlice = append(makeSlice, 100)
	otherMakeSlice := make([]int, 0, 10) // specifying length and capacity
	otherMakeSlice = append(otherMakeSlice, 99)
	fmt.Println(makeSlice)
	fmt.Println(otherMakeSlice)

	// slicing slices
	daSlice := []int{1, 2, 3, 4, 5}
	daSlice2 := daSlice[:2]
	daSlice3 := daSlice[1:]
	daSlice4 := daSlice[1:3]
	daSlice5 := daSlice[:]
	fmt.Println(daSlice)
	fmt.Println(daSlice2)
	fmt.Println(daSlice3)
	fmt.Println(daSlice4)
	fmt.Println(daSlice5)

	// slices share memory
	parentSlice := []int{100, 200, 300, 400}
	childSlice := parentSlice[:2]
	childSlice[1] = 99
	parentSlice[0] = 11
	fmt.Println(parentSlice)
	fmt.Println(childSlice)

	// wierd slice behavior
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// using copy()
	xForCopy := []int{1, 2, 3, 4}
	yForCopy := make([]int, 4)
	num := copy(yForCopy, xForCopy)
	fmt.Println(yForCopy, num)

	/*
		- Strings, Runes, and Bytes -
	*/
	var s string = "hello"
	var b byte = s[0]
	fmt.Println(b)

	var r rune = 'x'
	var s2 string = string(r)
	var b2 byte = 'y'
	var s3 string = string(b)
	fmt.Println(s2, b2, s3)

	/*
		- Maps -
	*/
	var nilMap map[string]int
	fmt.Println(nilMap)

	totalWins := map[string]int{}
	totalWins["usa"] = 2
	totalWins["france"] = 5
	fmt.Println(totalWins)
	fmt.Println(totalWins["france"])

	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	// comma - ok idiom
	m := map[string]int{
		"a": 5,
		"b": 3,
	}
	v, ok := m["a"]
	fmt.Println(v, ok)
	v, ok = m["c"]
	fmt.Println(v, ok)

	// deleting from maps
	delete(m, "a")
	fmt.Println(m)

	/*
		- Sets (just use a map)
	*/
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	/*
		- Structs -
	*/
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	fmt.Println(fred)

	bob := person{
		"bob",
		20,
		"sparky",
	}
	fmt.Println(bob)

	jim := person{
		age:  54,
		name: "jim",
		pet:  "fluffy",
	}
	fmt.Println(jim)
	jim.age = 99
	fmt.Println(jim.age)

	// anonymous structs
	var guy struct {
		name string
		age  int
		job  string
	}
	guy.name = "fred"
	guy.age = 89
	guy.job = "CEO"
	fmt.Println(guy)

	pet := struct {
		name string
		kind string
	}{
		name: "spike",
		kind: "dog",
	}
	fmt.Println(pet)
}
