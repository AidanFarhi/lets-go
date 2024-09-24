package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		inGoroutine := 1
		// every write to an unbuffered channel will cause
		// the writing goroutine to pause until another
		// goroutine reads from the same channel.
		// this goroutine cannot proceed until ch1 is read
		ch1 <- inGoroutine
		fromMain := <-ch2
		fmt.Println("goroutine:", inGoroutine, fromMain)
	}()
	inMain := 2
	// every write to an unbuffered channel will cause
	// the writing goroutine to pause until another
	// goroutine reads from the same channel.
	// this goroutine cannot proceed until ch1 is read
	ch2 <- inMain
	fromGoroutine := <-ch1
	fmt.Println("main:", inMain, fromGoroutine)
}
