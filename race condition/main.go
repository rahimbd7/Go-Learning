package main

import (
	"fmt"
	"sync"
)

var count int
var wg sync.WaitGroup
var mu sync.Mutex

func increment() {
	mu.Lock()
	count++
	mu.Unlock()
	// wg.Done()
}

func main() {
	for range 1000 {
		wg.Go(increment)
	}
	wg.Wait()
	fmt.Println("Final count:", count)
}
