package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 100
	randomNums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		randomNums = append(randomNums, rand.Int())
	}
	fmt.Println(randomNums[:5])
}
