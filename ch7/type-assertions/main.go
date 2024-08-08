package main

import (
	"fmt"
	"io"
)

type MyInt int

// a type switch
func checkType(i any) {
	switch j := i.(type) {
	case nil:
		fmt.Println("i is nil, j is any", j)
	case int:
		fmt.Println("j is of type int", j)
	case MyInt:
		fmt.Println("j is of type MyInt", j)
	case io.Reader:
		fmt.Println("j is of type io.Reader", j)
	case string:
		fmt.Println("j is of type string", j)
	case bool, rune:
		fmt.Println("i is either a bool or rune, so j is of type any", j)
	default:
		fmt.Println("no idea what i is, so j is of type any", j)
	}

}

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	// this is a type assertion.
	// type assertions can only be used
	// with interface types.
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)
	// this will cause a panic
	// i3 := i.(string)
	i4, ok := i.(int)
	if !ok {
		fmt.Errorf("unexpected type for %v", i)
	}
	fmt.Println(i4 + 1)

	var x, y, z any
	x = 20
	y = "hello"
	z = []string{"1"}
	checkType(x)
	checkType(y)
	checkType(z)
	checkType("yolo")

}
