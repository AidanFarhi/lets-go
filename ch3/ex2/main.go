package main

import (
	"fmt"
)

func main() {
	message := "Hi  and "
	runeSlice := []rune(message)
	fmt.Println(string(runeSlice[3]))
}
