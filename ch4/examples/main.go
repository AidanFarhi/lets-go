package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Shadowing Variables
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	// All go keywords are defined in the universe block,
	// which is the block that contains all other blocks.
	// They can be shadowed and reassigned.

	// fmt.Println(true)
	// true := 100
	// fmt.Println(true)

	// The if statement
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("too low:", n)
	} else if n > 5 {
		fmt.Println("too large:", n)
	} else {
		fmt.Println("perfect:", n)
	}

	// Scoping a variable to an if statement
	if n := rand.Intn(10); n == 0 {
		fmt.Println("too low:", n)
	} else if n > 5 {
		fmt.Println("too large:", n)
	} else {
		fmt.Println("perfect:", n)
	}

	// for loops, four ways

	// the complete for statement
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// the condition only for statement
	i := 1
	for i < 100 {
		fmt.Println(i)
		i *= 3
	}

	// the infinite for statement
	i = 1
	for {
		if i > 100 {
			break
		}
		fmt.Println(i)
		i *= 3
	}

	// the for-range statement
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	oddVals := []int{1, 3, 5, 7, 9}
	for _, v := range oddVals {
		fmt.Println(v)
	}

	nameSet := map[string]bool{"bob": true, "joe": true}
	for k := range nameSet {
		fmt.Println(k)
	}

	// continue
	for i := 1; i <= 10; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
