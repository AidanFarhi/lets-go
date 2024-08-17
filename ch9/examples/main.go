package main

import (
	"errors"
	"fmt"
	"os"
)

// you can use errors.New() to create a new error from a string
func DivMod(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return num / denom, num % denom, nil
}

// you can use fmt.Errorf() for formatting the error string
func DoubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d is not an even number", i)
	}
	return i * 2, nil
}

func main() {

	x := 20
	y := 3
	rem, mod, err := DivMod(x, y)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(rem, mod)

	res, err := DoubleEven(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
