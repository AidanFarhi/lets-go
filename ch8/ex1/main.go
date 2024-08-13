package main

import "fmt"

/*
Write a generic function that doubles the value of any integer or
float that’s passed in to it. Define any needed generic interfaces.
*/

type Doubleable interface {
	int | int16 | int32 | int64 | uint | uint16 | uint32 |
		uint64 | uintptr | float32 | float64
}

func Double[T Doubleable](x T) T {
	return x * 2
}

func main() {
	i1, i2 := 10, 20
	f1, f2 := 13.2, 53.66
	fmt.Println(Double(i1))
	fmt.Println(Double(i2))
	fmt.Println(Double(f1))
	fmt.Println(Double(f2))
}
