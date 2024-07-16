package main

import "fmt"

func main() {
	s := "Hi 👩 and 👨"
	r := []rune(s)
	fmt.Println(string(r[3]))
}
