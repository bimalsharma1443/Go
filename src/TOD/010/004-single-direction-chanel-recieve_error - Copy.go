package main

import "fmt"

func main() {

	cr := make(<-chan int)

	go func() {
		cr <- 42
	}()

	fmt.Println(<-cr)
	fmt.Printf("--------------\n")
	fmt.Printf("send chanel type \t%T\n", cr)
}
