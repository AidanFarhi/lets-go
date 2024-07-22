package main

import (
	"fmt"
	"math/rand"
)

/*
Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
*/

func main() {
	const n = 100
	randomNums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		randomNums = append(randomNums, rand.Intn(n+1))
	}
	fmt.Println(randomNums)
}
