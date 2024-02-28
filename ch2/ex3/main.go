package main

import "fmt"

/*
	Write a program with three variables,
	one named b of type byte, one named smallI of type int32,
	and one named bigI of type uint64. Assign each variable
	the maximum legal value for its type; then add 1 to each
	variable. Print out their values.
*/

func main() {

	var b byte
	var smallI int32
	var bigI uint64

	b = 127
	smallI = 2147483647
	bigI = 18446744073709551615

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
