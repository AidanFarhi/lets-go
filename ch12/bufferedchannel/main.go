package main

import (
	"fmt"
	"math/rand"
	"time"
)

func process(val int) int {
	n := rand.Intn(4)
	time.Sleep(time.Duration(n) * time.Second)
	return val * 2
}

func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func main() {
	data := make(chan int, 10)
	for i := 0; i < 10; i++ {
		data <- i
	}
	result := processChannel(data)
	fmt.Println(result)
}
