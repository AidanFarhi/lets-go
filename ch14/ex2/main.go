package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
Write a program that adds randomly generated numbers between 0 (inclusive) and 100,000,000
(exclusive) together until one of two things happen: the number 1234 is generated or 2
seconds has passed. Print out the sum, the number of iterations, and the reason for
ending (timeout or number reached).
*/

func generateNums(ctx context.Context) (int, int, error) {
	sum := 0
	count := 0
	for {
		select {
		case <-ctx.Done():
			return sum, count, ctx.Err()
		default:
			n := rand.Intn(100_000_000)
			if n == 1_234 {
				return sum, count, nil
			}
			sum += n
			count++
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	sum, count, err := generateNums(ctx)
	if err != nil {
		fmt.Println("sum:", sum, "number of iterations:", count, "did not find 1234")
	} else {
		fmt.Println("sum:", sum, "number of iterations:", count, "found 1234")
	}
}
