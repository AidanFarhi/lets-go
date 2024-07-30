package main

import "fmt"

/*
Start a new program. In main, declare an int variable called total.
Write a for loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive).

The body of the for loop should be as follows:

total := total + i
fmt.Println(total)

After the for loop, print out the value of total.
What is printed out? What is the likely bug in this code?
*/

func main() {
	total := 0
	for i := 0; i < 10; i++ {
		total := total + i // this is shadowing the outer total variable. therefore, it is not being updated
		fmt.Println(total)
	}
	fmt.Println(total)
}
