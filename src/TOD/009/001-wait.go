package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	bar()
	go too()
	go loo()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func too() {
	for i := 0; i < 10; i++ {
		fmt.Println("too:", i)
	}
}

func loo() {
	for i := 0; i < 10; i++ {
		fmt.Println("loo:", i)
	}
	wg.Done()
}
