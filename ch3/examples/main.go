package main

import (
	"fmt"
	"maps"
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

	var lenCapSlice []int
	fmt.Println(lenCapSlice, len(lenCapSlice), cap(lenCapSlice))
	lenCapSlice = append(lenCapSlice, 10)
	fmt.Println(lenCapSlice, len(lenCapSlice), cap(lenCapSlice))
	lenCapSlice = append(lenCapSlice, 20)
	fmt.Println(lenCapSlice, len(lenCapSlice), cap(lenCapSlice))
	lenCapSlice = append(lenCapSlice, 30)
	fmt.Println(lenCapSlice, len(lenCapSlice), cap(lenCapSlice))
	lenCapSlice = append(lenCapSlice, 40)
	fmt.Println(lenCapSlice, len(lenCapSlice), cap(lenCapSlice))
	lenCapSlice = append(lenCapSlice, 50)
	fmt.Println(lenCapSlice, len(lenCapSlice), cap(lenCapSlice))

	// using make()
	mSlice := make([]int, 5)    // len: 5 cap: 5
	mSlice = append(mSlice, 10) // this will be appended to the end of the slice

	fmt.Println(mSlice)

	mSlice2 := make([]int, 0, 10)  // len: 0 cap: 10
	mSlice2 = append(mSlice2, 100) // this will be the first value in the slice

	fmt.Println(mSlice2)

	// emptying a slice
	stSlice := []string{"one", "two", "three"}
	fmt.Println(stSlice, len(stSlice))
	clear(stSlice)                     // this will set all elements to their zero value
	fmt.Println(stSlice, len(stSlice)) // the length remains the same

	// slicing slices
	slSlice := []string{"a", "b", "c", "d", "e"}
	sl1 := slSlice[:2]
	sl2 := slSlice[1:]
	sl3 := slSlice[1:3]

	fmt.Println(sl1, sl2, sl3)

	// slices share memory
	sl1[1] = "y"
	sl2[0] = "x"

	fmt.Println(sl1, sl2, sl3)

	// append with slices is a bad idea
	appSlice := []string{"a", "b", "c", "d"}
	appSlice2 := appSlice[:2]
	fmt.Println(appSlice, len(appSlice), cap(appSlice))
	fmt.Println(appSlice2, len(appSlice2), cap(appSlice2))
	appSlice2 = append(appSlice2, "z")
	fmt.Println(appSlice, len(appSlice), cap(appSlice))
	fmt.Println(appSlice2, len(appSlice2), cap(appSlice2))

	// brain twister
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// x = [a, b, i, j, y]
	// y = [a, b, i, j, y]
	// z = [i, j, y]

	// using a full slice expression will prevent issues
	x = make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y = x[:2:2]
	z = x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println(cap(x), cap(y), cap(z))
	// x = [a, b, c, d, x]
	// y = [a, b, i, j, k]
	// z = [c, d, y]

	// using copy()
	x = make([]string, 0, 4)
	x = append(x, "1", "2", "3", "4")
	y = make([]string, 4)
	num := copy(y, x)
	fmt.Println(y, num)

	// you can also use copy with arrays by taking a slice of an array
	a := [4]int{1, 2, 3, 4}
	s := []int{5, 6, 7, 8}
	c := make([]int, 2)
	copy(c, a[:])
	fmt.Println(c)
	copy(a[:], s)
	fmt.Println(a)

	// convert arrays to slices
	a = [4]int{1, 2, 3, 4}
	aConv := a[:]
	aConv = append(aConv, 100)
	fmt.Println(aConv)

	// converting slices to arrays
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 100
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)

	// convert slice to a pointer to an array. their memory will be shared
	xArrayPointer := (*[4]int)(xSlice)
	xSlice[2] = 99
	xArrayPointer[0] = 88
	fmt.Println(xArrayPointer)
	fmt.Println(xSlice)

	//---------- Strings, Runes, and Bytes -----------

	// A string is a sequence of bytes. Go libraries assume
	// that a string is UTF-8 encoded.
	// A UTF-8 code point is 1-4 bytes long.

	// you can index and get slices from strings
	str := "hello there"
	b := str[3]
	strSlice := str[:3]

	fmt.Println(b, strSlice)

	// len() returns the length in bytes, not the
	// number of UTF-8 code points.
	fmt.Println(len(str))

	// Runes and Bytes can be converted to strings.
	r := 'x'
	rStr := string(r)
	bStr := 'y'
	rStr2 := string(bStr)

	fmt.Println(rStr, rStr2)

	// Strings can be converted to slices of bytes or runes
	byteSlice := []byte(str)
	runeSlice := []rune(str)

	fmt.Println(byteSlice)
	fmt.Println(runeSlice)

	//-------- Maps ----------
	var nilMap map[string]int // nil map variable

	fmt.Println(nilMap["one"])
	fmt.Println(nilMap)
	// nilMap["one"] = 1 this will cause a panic

	totalWins := map[string]int{} // empty map literal

	fmt.Println(totalWins["teamOne"])
	fmt.Println(totalWins)

	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	totalWins["Kittens"]++

	fmt.Println(totalWins)

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	fmt.Println(teams["Orcas"])

	ages := make(map[string]int, 10)
	ages["Bob"] = 22

	fmt.Println(ages)

	m := map[string]int{
		"hello": 5,
		"world": 0,
		"food":  100,
	}

	v, ok := m["hello"]

	fmt.Println(v, ok)

	v, ok = m["world"]

	fmt.Println(v, ok)

	v, ok = m["goodbye"]

	fmt.Println(v, ok)

	delete(m, "hello")
	v, ok = m["hello"]

	fmt.Println(v, ok)

	clear(m)
	v, ok = m["food"]

	fmt.Println(v, ok)
	fmt.Println(len(m))

	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	m2 := map[string]int{
		"two": 2,
		"one": 1,
	}

	fmt.Println(maps.Equal(m1, m2))

	// implementing a set using bool
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[100])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	// implementing a set using struct
	stringSet := map[string]struct{}{}
	strVals := []string{"a", "b", "a", "c", "d", "b", "a", "g"}
	for _, v := range strVals {
		stringSet[v] = struct{}{}
	}

	if _, ok := stringSet["a"]; ok {
		fmt.Println("a is in the set")
	}

	//------------- Structs --------------
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	bob := person{}
	julia := person{
		"Julia",
		40,
		"cat",
	}
	joe := person{
		age:  54,
		name: "joe",
	}

	fmt.Println(fred)
	fmt.Println(bob)
	fmt.Println(julia)
	fmt.Println(joe)
}
