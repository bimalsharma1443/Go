package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops int64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&ops, 1)
			}
			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Println("ops : ", ops)

}
