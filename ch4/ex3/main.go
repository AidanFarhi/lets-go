package main

import "fmt"

func main() {
	total := 0
	for i := 0; i < 10; i++ {
		total := total + i // this is shadowing the outer variable
		fmt.Println(total)
	}
	fmt.Println(total) // this will print 0
}
