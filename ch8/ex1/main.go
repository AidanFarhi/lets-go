package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint |
		uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Double[T Number](number T) T {
	return number * 2
}

func main() {
	fmt.Println(Double(3))
	fmt.Println(Double(23.534))
}
