package main

import (
	"fmt"
)

/*
Define a generic interface called Printable that matches a
type that implements fmt.Stringer and has an underlying
type of int or float64. Define types that meet this interface.
Write a function that takes in a Printable and prints its value
to the screen using fmt.Println.
*/

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintableInt int

func (pi PrintableInt) String() string {
	return fmt.Sprintf("%d", pi)
}

type PrintableFloat float64

func (pf PrintableFloat) String() string {
	return fmt.Sprintf("%.2f", pf)
}

func PrintPrintable[T Printable](printable T) {
	fmt.Println(printable)
}

func main() {
	var x PrintableInt = 100
	var y PrintableFloat = 32.415
	PrintPrintable(x)
	PrintPrintable(y)
}
