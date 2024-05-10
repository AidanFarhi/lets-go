package main

import (
	"fmt"
)

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
	return fmt.Sprintf("%f", pf)
}

func PrintItem[T Printable](p T) {
	fmt.Println(p)
}

func main() {
	num := PrintableInt(10)
	num2 := PrintableFloat(23.516)
	PrintItem(num)
	PrintItem(num2)
}
