package main

import (
	"fmt"
	"sync"
)

/*
Create a function that launches three goroutines that communicate using a
channel. The first two goroutines each write 10 numbers to the channel.
The third goroutine reads all the numbers from the channel and prints
them out. The function should exit when all values have been printed out.
Make sure that none of the goroutines leak. You can create additional
goroutines if needed.
*/
func process() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for n := range ch {
			fmt.Println(n)
		}
	}()
	wg2.Wait()
}

func main() {
	process()
}
