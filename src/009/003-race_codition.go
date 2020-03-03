package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("num of cpu ", runtime.NumCPU())
	fmt.Println("Goroutines ", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			//time.sleep(time.second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println(counter)
	}

	fmt.Println("counter : ", counter)

	fmt.Println("Goroutines ", runtime.NumGoroutine())
}
