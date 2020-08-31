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
	var mutex sync.Mutex

	for i := 0; i < 5000; i++ {
		wg.Add(1)

		go func() {
			mutex.Lock()
			for j := 0; j < 10; j++ {
				counter++
			}
			wg.Done()
			mutex.Unlock()
		}()

	}

	wg.Wait()

	fmt.Println("counter : ", counter)
	fmt.Println("Goroutines ", runtime.NumGoroutine())

}
