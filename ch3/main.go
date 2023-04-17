package main

import "fmt"

func main() {
	// arrays
	// in go, the size of the array is part of the type
	var x [3]int
	fmt.Println(x)

	var y = [3]int{1, 2, 3}
	fmt.Println(y)

	var z = [...]int{100, 200, 300}
	fmt.Println(z)

	fmt.Println(y == z)
	fmt.Println(len(z))

	// slices
	// in go, the length of a slice is not part of the type
	var mySlice = []int{1, 2, 3}
	fmt.Println(mySlice)

	var emptySlice []int
	fmt.Println(emptySlice)
	fmt.Println(emptySlice == nil)

	var newSlice []int
	newSlice = append(newSlice, 11)
	newSlice = append(newSlice, 22)
	fmt.Println(newSlice)

	var sliceOne = []int{55, 66, 77}
	var sliceTwo = []int{77, 88, 99}
	var sliceThree = append(sliceOne, sliceTwo...)
	fmt.Println(sliceThree)

	var growingSlice []int
	for i := 0; i < 10; i++ {
		fmt.Println(growingSlice, len(growingSlice), cap(growingSlice))
		growingSlice = append(growingSlice, i)
	}
}
