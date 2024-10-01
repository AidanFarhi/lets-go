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
			process(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
