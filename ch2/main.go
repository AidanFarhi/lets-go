package main

import (
	"fmt"
	"math/cmplx"
)

const (
	idKey   = "id"
	nameKey = "name"
)

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

	// complex numbers
	complexVal1 := complex(2.5, 3.1)
	complexVal2 := complex(10.2, 2)
	fmt.Println(complexVal1 + complexVal2)
	fmt.Println(complexVal1 - complexVal2)
	fmt.Println(complexVal1 * complexVal2)
	fmt.Println(complexVal1 / complexVal2)
	fmt.Println(real(complexVal1))
	fmt.Println(imag(complexVal1))
	fmt.Println(cmplx.Abs(complexVal1))

	// type conversions
	var type1 int = 10
	var type2 float64 = 30.2
	var type3 float64 = float64(type1) + type2
	var type4 int = type1 + int(type2)
	fmt.Println(type3, type4)

	// var vs. :=
	var myVal = 10
	myVal2 := 10
	fmt.Println(myVal, myVal2)

	numVal, stringVal := 100, "hello"
	fmt.Println(numVal, stringVal)

	// using const
	const myConst = "yolo"
	fmt.Println(myConst)
	newId := idKey + "123"
	fmt.Println(newId)
	nameKey := "newNameKey"
	fmt.Println(nameKey)

	// untyped constant
	const untypedConstant = 10
	var val1 int = untypedConstant
	var val2 float64 = untypedConstant
	var val3 byte = untypedConstant
	fmt.Println(val1, val2, val3)
}
