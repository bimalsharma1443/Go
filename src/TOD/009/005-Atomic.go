package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("num of cpu ", runtime.NumCPU())
	fmt.Println("Goroutines ", runtime.NumGoroutine())

	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines ", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("counter end : ", counter)

	fmt.Println("Goroutines ", runtime.NumGoroutine())
}
