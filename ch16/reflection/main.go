package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A string
	B int
}

func main() {
	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name())
	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name())
}
