package main

import "fmt"

func main() {
	fmt.Println("3 +4 =", add(3, 4))
	fmt.Println("4 +4 =", add(4, 4))
	fmt.Println("6 +4 =", add(6, 4))
}

func add(i ...int) int {
	x := 0
	for _, v := range i {
		x += v
	}
	return x
}
