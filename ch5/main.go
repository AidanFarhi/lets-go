package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) {
	if opts.FirstName != "" {
		fmt.Println("FirstName", opts.FirstName)
	}
	if opts.LastName != "" {
		fmt.Println("LastName", opts.LastName)
	}
	fmt.Println("Age", opts.Age)
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func multBy2And3(x int) (twoResult int, threeResult int) {
	twoResult = x * 2
	threeResult = x * 3
	return twoResult, threeResult
}

func add(x, y int) int      { return x + y }
func subtract(x, y int) int { return x - y }
func multipy(x, y int) int  { return x * y }
func divide(x, y int) int   { return x / y }

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func makeMultiplier(base int) func(int) int {
	return func(factor int) int { return base * factor }
}

func main() {
	/*
		- Functions -
	*/
	result := div(5, 2)
	fmt.Println(result)

	// simulate named parameters
	MyFunc(MyFuncOpts{
		LastName: "Smith",
	})
	MyFunc(MyFuncOpts{
		FirstName: "Bob",
		Age:       20,
	})

	// variadic parameters
	base := 10
	a := []int{7, 6}
	r1 := addTo(base, 1, 2, 3)
	r2 := addTo(base, 9, 8, 7, 6, 5)
	r3 := addTo(base, a...)
	r4 := addTo(base, []int{11, 22, 33}...)
	fmt.Println(r1, r2, r3, r4)

	// multiple return values
	res, rem, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res, rem)

	// named return values
	twoRes, threeRes := multBy2And3(5)
	fmt.Println(twoRes, threeRes)

	// functions are values
	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": subtract,
		"*": multipy,
		"/": divide,
	}
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"4", "/", "2"},
	}
	for _, expression := range expressions {
		p1, _ := strconv.Atoi(expression[0])
		op := expression[1]
		p2, _ := strconv.Atoi(expression[2])
		opFunc := opMap[op]
		opResult := opFunc(p1, p2)
		fmt.Println(p1, op, p2, " =", opResult)
	}

	// anonymous functions
	for i := 0; i < 3; i++ {
		func(x int) {
			fmt.Println("printing", x, "from anon func")
		}(i)
	}

	/*
		- Closures -
		Functions declared within functions are able to access
		and modify variables declared in the outer function
	*/

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"Bob", "Smith", 45},
		{"Jim", "Johnson", 22},
		{"Alvin", "Bledsoe", 76},
	}
	fmt.Println(people)
	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
	// sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	twoMultiplier := makeMultiplier(2)
	threeMultiplier := makeMultiplier(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoMultiplier(i), threeMultiplier(i))
	}
}
