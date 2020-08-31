package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	go func() {
		fmt.Println("12")
		time.Sleep(1 * time.Second)
		c1 <- "result 1"
	}()

	fmt.Println("17")

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 1")
	}

	fmt.Println("26")

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
