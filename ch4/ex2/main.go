package main

import (
	"fmt"
	"math/rand"
)

/*
Loop over the slice you created in exercise 1.
For each value in the slice, apply the following rules:

If the value is divisible by 2, print “Two!”

If the value is divisible by 3, print “Three!”

If the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.

Otherwise, print “Never mind”.
*/

func main() {

	const n = 100
	randomNums := make([]int, 0, 100)
	for i := 0; i < n; i++ {
		randomNums = append(randomNums, rand.Intn(n+1))
	}

	for _, num := range randomNums {
		switch {
		case num%6 == 0:
			fmt.Println(num, "Six!")
		case num%2 == 0:
			fmt.Println(num, "Two!")
		case num%3 == 0:
			fmt.Println(num, "Three!")
		default:
			fmt.Println(num, "Never mind")
		}
	}
}
