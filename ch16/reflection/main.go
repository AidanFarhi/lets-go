package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int    `myTag:"value"`
	B string `myTag:"value2"`
}

func changeInt(i *int) {
	*i = 100
}

func changeIntReflect(i *int) {
	iv := reflect.ValueOf(i)
	iv.Elem().SetInt(100)
}

func hasNoValue(i any) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Func,
		reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}

func main() {
	var x int
	f := Foo{}

	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name())

	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name())

	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())

	xk := reflect.TypeOf(x).Kind()
	fmt.Println(xk)

	for i := 0; i < ft.NumField(); i++ {
		currf := ft.Field(i)
		fmt.Println(currf.Name, currf.Type.Name(), currf.Tag.Get("myTag"))
	}

	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)
	s2 := sv.Interface().([]string)
	for _, el := range s2 {
		fmt.Println(el)
	}

	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	ivv.SetInt(100)
	fmt.Println(i)

	j := 20
	changeInt(&j)
	fmt.Println(j)

	k := 30
	changeIntReflect(&k)
	fmt.Println(k)

	var m map[int]int
	m2 := map[int]int{}
	fmt.Println(hasNoValue(k))
	fmt.Println(hasNoValue(m))
	fmt.Println(hasNoValue(m2))
}
