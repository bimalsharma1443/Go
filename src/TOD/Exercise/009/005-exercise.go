package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	var bg = 100

	var wg sync.WaitGroup

	wg.Add(bg)

	for i := 0; i < bg; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter : ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter final value : ", counter)

}
