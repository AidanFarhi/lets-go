package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func process(id int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(6)))
	fmt.Printf("%d finished\n", id)
}

func main() {
	nProc := 20
	var wg sync.WaitGroup
	wg.Add(nProc)
	for i := 0; i < nProc; i++ {
		go func() {
			// this ensures that wg.Done() is called even
			// if process() causes a panic
			defer wg.Done()
			process(i)
		}()
	}
	wg.Wait()
}
