package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Write a function that builds a map[int]float64 where the keys are the
numbers from 0 (inclusive) to 100,000 (exclusive) and the values are
the square roots of those numbers (use the math.Sqrt function to
calculate square roots). Use sync.OnceValue to generate a function that caches
the map returned by this function and use the cached value to look up square
roots for every 1,000th number from 0 to 100,000.
*/

func initRootMap() map[int]float64 {
	fmt.Println("initRootMap called")
	rootMap := make(map[int]float64)
	for i := 0; i < 100_000; i++ {
		rootMap[i] = math.Sqrt(float64(i))
	}
	return rootMap
}

var getCachedRootMap func() map[int]float64 = sync.OnceValue(initRootMap)

func main() {
	valsToCheck := []int{4, 16, 25, 100}
	for _, v := range valsToCheck {
		vSquareRoot := getCachedRootMap()[v]
		fmt.Printf("Square root of: %d = %.2f\n", v, vSquareRoot)
	}
}
