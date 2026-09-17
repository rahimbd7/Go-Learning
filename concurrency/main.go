package main

import (
	"fmt"
	"sync"
	"time"
)

func task1() {
	fmt.Println("task 1 loading")
	time.Sleep(3 * time.Second)
	fmt.Println("task 1 done")
	// wg.Done()
}
func task2() {
	fmt.Println("task 2 loading")
	time.Sleep(1 * time.Second)
	fmt.Println("task 2 done")
	// wg.Done()
}
func task3() {
	fmt.Println("task 3 loading")
	time.Sleep(2 * time.Second)
	fmt.Println("task 3 done")
	// wg.Done()
}

var wg sync.WaitGroup

func main() {

	start := time.Now()
	// wg.Add(1) // Add a count of 1 to the WaitGroup for task1
	// go task1()

	// wg.Add(1) // Add a count of 1 to the WaitGroup for task2
	// go task2()

	// wg.Add(1) // Add a count of 1 to the WaitGroup for task3
	// go task3()


	wg.Go(task1)
	wg.Go(task2)
	wg.Go(task3)

	wg.Wait()
	fmt.Println("All tasks done in", time.Since(start))
}
