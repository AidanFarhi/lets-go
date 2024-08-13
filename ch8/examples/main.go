package main

import (
	"errors"
	"fmt"
	"math"
)

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// This works with slices of any type.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s for
// which f returned true.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair2.Val1.Diff(pair2.Val2)
	if d1 < d2 {
		return pair1
	}
	return pair2
}

type Point2D struct {
	X, Y int
}

func (p2 Point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p2.X, p2.Y)
}

func (p2 Point2D) Diff(from Point2D) float64 {
	x := p2.X - from.X
	y := p2.Y - from.Y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

type Point3D struct {
	X, Y, Z int
}

func (p3 Point3D) String() string {
	return fmt.Sprintf("{%d,%d,%d}", p3.X, p3.Y, p3.Z)
}

func (p3 Point3D) Diff(from Point3D) float64 {
	x := p3.X - from.X
	y := p3.Y - from.Y
	z := p3.Z - from.Z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z))
}

// define a type element using type terms
type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

// the squiggly makes the type terms valid for
// any type that has the type term as its underlying type.
type AnyInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func DivAndRemainder[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func AnyDivAndRemainder[T AnyInteger](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

// type inference
type IntegerInf interface {
	int | int8 | int16 | int32 | int64 | uint |
		uint8 | uint16 | uint32 | uint64
}

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}

type Thinger interface {
	Thing()
}

type ThingerInt int

func (t ThingerInt) Thing() {
	fmt.Println("ThingIng:", t)
}

type ThingerSlice []int

func (t ThingerSlice) Thing() {
	fmt.Println("ThingSlice:", t)
}

func Comparer[T comparable](t1, t2 T) {
	if t1 == t2 {
		fmt.Println("equal!")
	}
}

func main() {

	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Println(intStack.Contains(10))
	fmt.Println(intStack.Contains(50))
	v, ok := intStack.Pop()
	for ok {
		fmt.Println(v)
		v, ok = intStack.Pop()
	}

	words := []string{"One", "Potato", "Two", "Potato"}
	filtered := Filter(words, func(s string) bool {
		return s != "Potato"
	})
	fmt.Println(filtered)
	lengths := Map(filtered, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)
	sum := Reduce(lengths, 0, func(acc int, val int) int {
		return acc + val
	})
	fmt.Println(sum)

	pair2Da := Pair[Point2D]{Point2D{1, 1}, Point2D{5, 5}}
	pair2Db := Pair[Point2D]{Point2D{10, 10}, Point2D{15, 5}}
	closer := FindCloser(pair2Da, pair2Db)
	fmt.Println(closer)
	pair3Da := Pair[Point3D]{Point3D{1, 1, 10}, Point3D{5, 5, 0}}
	pair3Db := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{11, 5, 0}}
	closer2 := FindCloser(pair3Da, pair3Db)
	fmt.Println(closer2)

	var a uint = 18_446_744_073_709_551_615
	var b uint = 9_223_372_036_854_775_808
	fmt.Println(DivAndRemainder(a, b))

	// this will not compile. the type terms must match
	// exactly.
	// type MyInt int
	// var myA MyInt = 10
	// var myB MyInt = 20
	// fmt.Println(DivAndRemainder(myA, myB))

	// this will work
	type MyInt int
	var myA MyInt = 10
	var myB MyInt = 20
	fmt.Println(AnyDivAndRemainder(myA, myB))

	// type inference
	var num1 int = 10
	num2 := Convert[int, int64](num1)
	fmt.Println(num2)

	var a1 int = 10
	var a2 int = 10
	Comparer(a1, a2)

	var a3 ThingerInt = 20
	var a4 ThingerInt = 20
	Comparer(a3, a4)

	var a5 ThingerSlice = []int{1, 2, 3}
	var a6 ThingerSlice = []int{1, 2, 3}
	// this will not compile
	// Comparer(a5, a6)

	var a7 Thinger = a5
	var a8 Thinger = a6
	// this will compile, but panic at runtime
	// Comparer(a7, a8)
	fmt.Println(a7, a8)
}
