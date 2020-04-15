package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go foo()
	boo()
	wg.Wait()
}

func foo() {
	fmt.Println("I am in foo")
	wg.Done()
}

func boo() {
	fmt.Println("I am in boo")
}
