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
}
