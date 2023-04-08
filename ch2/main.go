package main

import "fmt"

func main() {

	// bools
	var flag bool // no value assigned, set to false
	var isAwesome = true
	fmt.Println(flag, isAwesome)

	// integer types in go
	var intVal1 int8 = 10
	var intVal2 int16 = 100
	var intVal3 int32 = 100
	var intVal4 int64 = 1_000
	var intVal5 uint8 = 10
	var intVal6 uint16 = 100
	var intVal7 uint32 = 100
	var intVal8 uint64 = 1_000
	fmt.Println(intVal1, intVal2, intVal3, intVal4)
	fmt.Println(intVal5, intVal6, intVal7, intVal8)

	// integer operations
	var x int = 10
	x *= 2
	fmt.Println(x)

	// dividing floats
	var floatVal float64 = 100
	fmt.Println(floatVal / 0) // +Inf
	floatVal = 0
	fmt.Println(floatVal / 0) // NaN
}
