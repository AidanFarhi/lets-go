package main

import "fmt"

func main() {

	// ---------------- Strings and Runes ----------------

	// Runes: run == int32
	var r1 rune = 'a'
	var r2 int32 = 'z'

	fmt.Println(r1, r2)

	// Interpreted string literals
	s1 := "hello world"
	s2 := "how are you"

	fmt.Println(s1, s2)

	// Raw string literals
	rs1 := `
	what do you mean?
	ok "I'm not sure why"...
	`
	rs2 := `
	<div>
		<p>Hello World\n</p>
	</div>
	`

	fmt.Println(rs1, rs2)

	// Booleans
	var flag bool // zero value for booleans is false
	var isAwsome = true

	fmt.Println(flag, isAwsome)

	// ---------------- Numeric Types ----------------

	// Signed ints
	var num1 int8 = 127
	var num2 int16 = 32_767
	var num3 int32 = 2_147_483_647
	var num4 int64 = 9_223_372_036_854_775_807

	fmt.Println(num1, num2, num3, num4)

	// Unsinged ints
	var unum1 uint8 = 255
	var unum2 uint16 = 65_535
	var unum3 uint32 = 4_294_967_295
	var unum4 uint64 = 18_446_744_073_709_551_615

	fmt.Println(unum1, unum2, unum3, unum4)

	// byte == uint8
	var bnum1 byte = 100
	var snum1 uint8 = 100

	fmt.Println(bnum1 == snum1)

	// Floating numbers
	var fnum1 float32 = 3.40282346638528859811704183484516925440e+38
	var fnum2 float64 = 1.797693134862315708145274237317043567981e+308

	fmt.Println(fnum1, fnum2)

	// Complex numbers
	var cNum = complex(20.3, 10.2)

	fmt.Println(real(cNum), imag(cNum), cNum)

	// ---------------- Type Conversions ----------------

	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)

	fmt.Println(sum1, sum2)

	// ---------------- Declaring Variables ----------------

	// using a declaration list
	var (
		xInt int
		yInt     = 20
		zInt int = 30
		d, e     = 40, "hello"
		f, g string
	)

	fmt.Println(xInt, yInt, zInt, d, e, f, g)

	// ---------------- Using const ----------------

	// There are only a few values that constants can hold:
	// - numeric literals
	// - true and false
	// - strings
	// - runes
	// - return values of complex(), real(), imag(), len(), and cap()
	// - expressions that consist of operators and the preceding values

	const cString = "hello"
	const cVal = 10
	// cString = "goodbye" this will not compile
	// cVal = cVal + 1     this will not compile

	// the below will not compile
	// type Person struct {
	// 	name string
	// }
	// const myPerson = Person{}

	const flexVal = 10
	var fVal1 float32 = flexVal
	var fVal2 byte = flexVal

	fmt.Println(cString, cVal, fVal1, fVal2)

}
