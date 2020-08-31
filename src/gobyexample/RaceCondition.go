package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	fmt.Println("num of cpu ", runtime.NumCPU())
	fmt.Println("Goroutines ", runtime.NumGoroutine())

	var wg sync.WaitGroup

	for i := 0; i < 5000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 10; j++ {
				counter++
			}
			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Println("counter : ", counter)
	fmt.Println("Goroutines ", runtime.NumGoroutine())

}
