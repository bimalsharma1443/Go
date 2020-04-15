package main

import "fmt"

func main() {
	func() {
		fmt.Println("Clouser function")
	}()

	incInt := call()

	fmt.Println(incInt())
	fmt.Println(incInt())
	fmt.Println(incInt())

	incInt1 := call()

	fmt.Println(incInt1())
	fmt.Println(incInt1())
	fmt.Println(incInt1())
}

func call() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
