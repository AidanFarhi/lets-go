package main

import (
	"fmt"
	format "package-example/do-format"
	"package-example/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
