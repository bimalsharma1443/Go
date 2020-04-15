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

	wg.Add(bg)

	for i := 0; i < bg; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			fmt.Println("counter : ", counter)
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter final value : ", counter)

}
