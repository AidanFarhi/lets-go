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
	*/
	fmt.Println("----------- Slices -----------")
}
