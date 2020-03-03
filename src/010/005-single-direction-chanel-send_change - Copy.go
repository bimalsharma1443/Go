package main

import "fmt"

func main() {

	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Printf("--------------\n")
	fmt.Printf("send chanel type \t%T\n", cr)
	fmt.Printf("send chanel type \t%T\n", cs)
	fmt.Printf("send chanel type \t%T\n", c)

	// not assigned from specific to general

	c = cr
	c = cs
}
