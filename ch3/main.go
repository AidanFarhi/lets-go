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

	originSlice := []int{1, 2, 3, 4, 5}
	childSliceOne := originSlice[:3]
	childSliceTwo := originSlice[1:4]
	fmt.Println(originSlice, childSliceOne, childSliceTwo)
	childSliceTwo[0] = 999
	fmt.Println(originSlice)

	sliceX := make([]int, 0, 5)
	sliceX = append(sliceX, 1, 2, 3, 4)
	sliceY := sliceX[:2]
	sliceZ := sliceX[2:]
	fmt.Println(sliceY)
	fmt.Println(sliceZ)
	fmt.Println(cap(sliceX), cap(sliceY), cap(sliceZ))
	sliceY = append(sliceY, 30, 40, 50)
	sliceX = append(sliceX, 60)
	fmt.Println(cap(sliceX))
	fmt.Println(cap(sliceZ))
	sliceZ = append(sliceZ, 70)
	fmt.Println("x:", sliceX)
	fmt.Println("y:", sliceY)
	fmt.Println("z:", sliceZ)

	// using copy
	copySlice := []int{1, 2, 3}
	copySlice2 := make([]int, 3)
	numElementsCopied := copy(copySlice2, copySlice)
	fmt.Println(copySlice2, numElementsCopied)

	sliceToCopy := []int{5, 6, 7, 8}
	arrToCopy := [...]int{1, 2, 3, 4}
	theNewSlice := make([]int, 2)
	copy(theNewSlice, arrToCopy[:])
	copy(arrToCopy[:], sliceToCopy)
	fmt.Println(sliceToCopy)
	fmt.Println(arrToCopy)
	fmt.Println(theNewSlice)

	// fun with strings
	s := "hello world"
	ch := s[3]
	fmt.Println(s, ch)

	// maps

}
