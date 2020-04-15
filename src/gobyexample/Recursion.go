package main

import "fmt"

func main() {

	var i uint64

	fmt.Println(" Please enter a number ")

	fmt.Scanln(&i)

	fmt.Printf("factorial of %v is %v", i, fact(i))
}

func fact(f uint64) uint64 {

	if f == 1 {
		return 1
	} else {
		return f * fact(f-1)
	}
}
