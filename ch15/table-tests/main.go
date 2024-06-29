package main

import (
	"fmt"
	"os"
	"table-tests/ops"
)

func main() {
	res, err := ops.DoMath(1, 2, "+")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(res)
}
