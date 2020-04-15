package main

import (
	"fmt"
	"runtime"
)

func main() {
	message := make(chan int)

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	go func() {
		for i := 1; i <= 10; i++ {
			message <- i
		}
		close(message)
	}()

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	for value := range message {
		fmt.Println(value)
	}
}
