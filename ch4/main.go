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

	// for-range loop
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenVals := []int{}
	for _, v := range vals {
		if v%2 == 0 {
			evenVals = append(evenVals, v)
		}
	}
	fmt.Println(evenVals)

	uniqueNames := map[string]bool{
		"Fred": false,
		"Raul": false,
		"Bob":  false,
	}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	// the val in a for-range loop is a copy
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)

	/*
		- Switch statement -
	*/
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			fmt.Println(word, "is exactly the right length:", size)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word")
		}
	}

loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not by 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	moreWords := []string{"hi", "salutations", "hello"}
	for _, word := range moreWords {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}

	/*
		- goto -
	*/
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}
