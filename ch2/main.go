package main

import "fmt"

func main() {

	/* Booleans */
	var flag bool // no value assigned, set to false
	fmt.Println(flag)

	/* Numeric Types */
	var int1 int8 = 127
	var int2 int16 = 32_767
	var int3 int32 = 2_147_483_647
	var int4 int64 = 9_223_372_036_854_775_807
	var int5 uint8 = 255
	var int6 uint16 = 65_535
	var int7 uint32 = 4_294_967_295
	var int8 uint64 = 18_446_744_073_709_551_615
	fmt.Println(int1, int2, int3, int4, int5, int6, int7, int8)
}
