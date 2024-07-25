package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// basic function
func Div(num, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

// simulating named parameters
type GreetFuncOps struct {
	Name string
	Age  int
}

func GreetFunc(opts GreetFuncOps) {
	fmt.Println("hello:", opts.Name)
	fmt.Println("you are:", opts.Age)
}

// variadic parameters
func AddTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// multiple return values
func DivAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

// named return values
func DivAndRemainderNamed(num, denom int) (result, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

// named return values confusion
func DivAndRemainderConfusing(num, denom int) (result, remainder int, err error) {
	result, remainder = 20, 30
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

// blank returns -- do not use
func DivAndRemainderBlankReturn(num, denom int) (result, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = num/denom, num%denom
	return
}

// functions are values
func GetLen(a string) int {
	return len(a)
}

func StringCalc(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func add(x, y int) int { return x + y }

func sub(x, y int) int { return x - y }

func mul(x, y int) int { return x * y }

func div(x, y int) int { return x / y }

// function type definition
type OpFunc func(int, int) int

// package scope anonymous functions
var (
	adder  = func(x, y int) int { return x + y }
	subber = func(x, y int) int { return x - y }
	multer = func(x, y int) int { return x * y }
	diver  = func(x, y int) int { return x / y }
)

// you can modify a package level variable
func ChangeAdder() {
	adder = func(x, y int) int { return x + x }
}

// returning functions from functions
type MultFunc func(int) int

func MultFuncGenerator(base int) MultFunc {
	return func(factor int) int {
		return base * factor
	}
}

// defer
func DeferExample() int {
	a := 10
	defer func(x int) {
		fmt.Println("first:", x)
	}(a)
	a = 20
	defer func(x int) {
		fmt.Println("second:", x)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}

// go is call by value
type person struct {
	age  int
	name string
}

func FailedUpdate(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "bob"
}

func ModMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func ModSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {

	result := Div(5, 2)
	fmt.Println(result)

	GreetFunc(GreetFuncOps{
		Name: "bob",
		Age:  23,
	})

	fmt.Println(AddTo(3))
	fmt.Println(AddTo(3, 4, 5, 6))
	fmt.Println(AddTo(100, 4, 5, 6))
	fmt.Println(AddTo(10, []int{1, 2, 3, 4}...))

	result, remainder, err := DivAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	result, remainder, err = DivAndRemainderNamed(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	result, remainder, err = DivAndRemainderConfusing(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	result, remainder, err = DivAndRemainderBlankReturn(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	// functions are values
	var stringFunc func(string) int
	// stringFunc("yolo") this would cause a panic
	stringFunc = GetLen
	result = stringFunc("yolo")
	fmt.Println(result)
	stringFunc = StringCalc
	result = stringFunc("yolo")
	fmt.Println(result)

	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	var opMapWithType = map[string]OpFunc{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("invalid expression:", exp)
			continue
		}
		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := exp[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(p1, op, p2, "=", result)
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("invalid expression:", exp)
			continue
		}
		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := exp[1]
		opFunc, ok := opMapWithType[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(p1, op, p2, "=", result)
	}

	// anonymous functions
	anon := func(j int) {
		fmt.Println("printing", j, "from inside anon")
	}
	for i := 0; i < 5; i++ {
		anon(i)
	}

	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside an anonymous function")
		}(i)
	}

	// using package scope defined anonymous functions
	nums := [][]int{{1, 2}, {3, 4}, {5, 6}}
	for _, pair := range nums {
		x, y := pair[0], pair[1]
		fmt.Println("adder:", adder(x, y))
		fmt.Println(subber(x, y))
		fmt.Println(multer(x, y))
		fmt.Println(diver(x, y))
	}

	// this will mutate the package defined adder anonymous function
	ChangeAdder()

	for _, pair := range nums {
		x, y := pair[0], pair[1]
		fmt.Println("adder:", adder(x, y))
		fmt.Println(subber(x, y))
		fmt.Println(multer(x, y))
		fmt.Println(diver(x, y))
	}

	// closures
	a := 20
	f := func() {
		fmt.Println("a in closure:", a)
		a = 100
	}
	f()
	fmt.Println("a after closure:", a)

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println("unsorted:", people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("sorted by age:", people)

	doubler := MultFuncGenerator(2)
	tripler := MultFuncGenerator(3)

	fmt.Println(doubler(2))
	fmt.Println(tripler(2))

	// defer
	DeferExample()

	// go is call by value
	p := person{}
	i := 2
	s := "Hello"
	FailedUpdate(i, s, p)
	fmt.Println(i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	ModMap(m)
	fmt.Println(m)

	sNums := []int{1, 2, 3}
	ModSlice(sNums)
	fmt.Println(sNums)
}
