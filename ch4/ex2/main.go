package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// generate a slice of random integers
	n := 100
	randomNums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		randomNums = append(randomNums, rand.Int())
	}

	// loop through and do fancy logic
	for _, num := range randomNums {
		switch {
		case num%2 == 0 && num%3 == 0:
			fmt.Println("Six!")
		case num%2 == 0:
			fmt.Println("Two!")
		case num%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
