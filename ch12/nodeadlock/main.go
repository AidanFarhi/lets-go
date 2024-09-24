package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1
		valFromCh2 := <-ch2
		fmt.Println("from ch2:", valFromCh2)
	}()
	select {
	case ch2 <- 2:
	case valFromCh1 := <-ch1:
		fmt.Println("from ch1:", valFromCh1)
	}
}
