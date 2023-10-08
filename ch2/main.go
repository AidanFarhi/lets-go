package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	/*
		In Go, all variables will be assigned a default zero value
	*/

	/* Strings */
	var stringVal string = "hello world"
	fmt.Println(stringVal)

	/* Booleans */
	var flag bool // no value assigned, set to false
	fmt.Println(flag)

	/* Numeric Types */
	var intVal1 int8 = 127
	var intVal2 int16 = 32_767
	var intVal3 int32 = 2_147_483_647
	var intVal4 int64 = 9_223_372_036_854_775_807
	var intVal5 uint8 = 255
	var intVal6 uint16 = 65_535
	var intVal7 uint32 = 4_294_967_295
	var intVal8 uint64 = 18_446_744_073_709_551_615
	fmt.Println(intVal1, intVal2, intVal3, intVal4, intVal5, intVal6, intVal7, intVal8)

	/* Integer Operators */
	var x int = 10
	x *= 2
	fmt.Println(x)
	fmt.Println(x > 10)

	/* Floating point types */
	var floatVal1 = 1232223234.00
	var floatVal2 = 12359172393482039482.2349234
	fmt.Println(floatVal1, floatVal2)
	fmt.Println(floatVal1 / 0) // +Inf

	/* Complex Numbers */
	var complexVal1 complex64 = complex(2.5, 3.1)
	var complexVal2 complex64 = complex(10.2, 2)
	fmt.Println(complexVal1 + complexVal2)
	fmt.Println(complexVal1 - complexVal2)
	fmt.Println(complexVal1 * complexVal2)
	fmt.Println(complexVal1 / complexVal2)
	fmt.Println(real(complexVal1))
	fmt.Println(imag(complexVal2))
	fmt.Println(cmplx.Abs(complex128(complexVal1)))

	/*
		- Type Conversions -
		Go does not allow automatic type conversion.
		You cannot treat treat another Go type as a boolean.

		Ex)

		// This is invalid Go
		var x string = ""
		if x {
			fm.Println("x is not empty")
		}
	*/
	var x1 int = 10
	var y1 float64 = 30.2
	var z1 float64 = float64(x1) + y1
	var d1 int = x1 + int(y1)
	fmt.Println(z1, d1)

	/* Variable Declaration */
	var foo int = 10
	var bar = 20 // defaults to int
	var baz int
	var bing, bang, boom int = 1, 2, 3
	var name, age = "bob", 25
	min := 0 // also defaults to int
	max := 100
	newMin, newMax := 0, 100
	weight, weightClass := 194, "heavyweight"
	var (
		address           = "123 street place"
		distanceMiles     = 424.32
		averageTravelTime = 24.3
	)
	fmt.Println(foo + bar)
	fmt.Println(baz)
	fmt.Println(bing, bang, boom)
	fmt.Println(name, age)
	fmt.Println(address, distanceMiles, averageTravelTime)
	fmt.Println(address, distanceMiles, averageTravelTime)
	fmt.Println(min, max)
	fmt.Println(newMin, newMax)
	fmt.Println(weight, weightClass)

	/*
		Constants in go can only be assigned the following:
		- Numeric literals (10, 10.232)
		- Strings ("hello", "world")
		- Runes
		- These built-in functions: complex, real, imag, len, and cap
		- Expressions that consist of operators and the preceding values

		In brief, constants are a way to give names to literals.
		There is no way in Go to declare that a variable is immutable.
	*/
	/* Using const */
	const x2 int64 = 10
	const (
		idKey   = "id"
		nameKey = "name"
	)
	const z = 20 * 10
	fmt.Println(x2)
	fmt.Println(idKey)
	fmt.Println(nameKey)
	fmt.Println(z)
	// x2 = x2 + 10 this will not compile
}
