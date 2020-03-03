package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	b := make(chan int, 2)

	go func() {
		c <- 42
		b <- 40
	}()

	fmt.Println(<-c)
	fmt.Println(<-b)
}
