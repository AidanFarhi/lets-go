package main

import (
	"fmt"
	"slices"
)

func main() {
	//------------- Arrays -------------

	// declaring arrays
	var aWithSize [3]int
	var aLiteral = [3]int{1, 2, 3}
	var aSparse = [12]int{1, 5: 4, 6, 10: 100, 15}
	var aLiteralAlternate = [...]int{100, 200, 300, 400, 500}

	fmt.Println(aWithSize, aLiteral, aSparse, aLiteralAlternate)

	// comparing arrays
	var a1 = [...]int{1, 1, 1}
	var a2 = [3]int{1, 1, 1}

	fmt.Println(a1 == a2)

	// multidimensional arrays
	var matrix [3][3]int

	fmt.Println(matrix)

	//------------- Slices -------------

	// declaring slices
	var sLiteral = []int{1, 2, 3}
	var sSparse = []int{1, 5: 4, 6, 10: 100, 15}
	var sMatrix = [][]int{{1, 2, 3}, {4, 5, 6}}
	var sNilSlice []int // the zero value for slices is nil

	fmt.Println(sLiteral, sSparse, sMatrix, sNilSlice)
	fmt.Println(sNilSlice == nil)
	fmt.Println(len(sSparse))

	// comparing slices using the "slices" package
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := []int{1, 2, 3}

	fmt.Println(slices.Equal(s1, s2))
	fmt.Println(slices.Equal(s1, s3))

	type Point struct {
		x int
		y int
	}

	pS1 := []Point{Point{1, 2}, Point{3, 4}}
	pS2 := []Point{Point{1, 2}, Point{3, 4}}
	pSEqual := slices.EqualFunc(pS1, pS2, func(p1, p2 Point) bool {
		return p1.x == p2.x && p1.y == p2.y
	})

	fmt.Println(pSEqual)

	// manipulating slices
	var appSl []int
	var appSl2 = []int{99, 88, 77}
	appSl = append(appSl, 10)
	appSl = append(appSl, 20)
	appSl = append(appSl, 100, 200, 300)
	appSl = append(appSl, appSl2...)

	fmt.Println(appSl)

	// length vs. capacity
	//
	// length: The number of consecutive memory locations that have been
	//         assigned a value.
	//
	// capacity: The number of consecutive memory locations reserved.

}
