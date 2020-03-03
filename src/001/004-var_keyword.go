package main

import "fmt"

var x = 40

var z = 45

func main() {

	y := 42

	fmt.Println("x is", x)
	fmt.Println("y is", y)

	foo()
}

func foo() {
	fmt.Println("z is", z)
}
