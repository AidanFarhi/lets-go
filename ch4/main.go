package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*
		- Blocks -
	*/

	// Shadowing Variables
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		x := 5
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10

	// Shadowing the built in names
	fmt.Println(true)
	true := 10
	fmt.Println(true)

	/*
		- if -
	*/
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low:", n)
	} else if n > 5 {
		fmt.Println("That's too high:", n)
	} else {
		fmt.Println("That's perfect:", n)
	}

	if x := rand.Intn(10); x == 0 {
		fmt.Println("That's too low:", x)
	} else if n > 5 {
		fmt.Println("That's too high:", x)
	} else {
		fmt.Println("That's perfect:", x)
	}

	/*
		- for loops -

		Go contains 4 different for loops.
			- C-style
			- Condition-only
			- Infinite
			- for-range
	*/

	// C-style for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		if i == 9 {
			fmt.Println()
		}
	}

	// Condition only for loop
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		if i == 9 {
			fmt.Println()
		}
		i++
	}

	// Inifinite for loop
	// for {
	// 	fmt.Println("hello")
	// }
}
