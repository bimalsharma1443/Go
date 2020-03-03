package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0
	var bg = 100

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(bg)

	for i := 0; i < bg; i++ {
		go func() {
			mutex.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("counter : ", counter)
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter final value : ", counter)

}
