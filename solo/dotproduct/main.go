package main

import "fmt"

func main() {
	a := [][]int{
		{1, 0},
		{0, 1},
	}
	b := [][]int{
		{1, 0},
		{0, 1},
	}
	// dp := [][]int{
	// 	{0, 0, 0},
	// 	{0, 0, 0},
	// 	{0, 0, 0},
	// }
	// dot product
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s := 0
			for k := 0; k < 3; k++ {
				p := a[i][j] * b[j][k]
				s += p
			}
			fmt.Println(s)
		}
	}
	// for _, r := range dp {
	// 	fmt.Println(r)
	// }
}
