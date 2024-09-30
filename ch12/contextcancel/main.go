package main

import (
	"context"
	"fmt"
)

func countTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- i
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := countTo(ctx, 10)
	for i := range ch {
		fmt.Println(i)
		if i > 5 {
			break
		}
	}
}
