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
	// 	{0, 0},
	// 	{0, 0},
	// }
	// dot product
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				aElement := a[i][j]
				bElement := b[j][k]
				fmt.Println("a", aElement)
				fmt.Println("b", bElement)
			}
		}
	}
	// for _, r := range dp {
	// 	fmt.Println(r)
	// }
}
