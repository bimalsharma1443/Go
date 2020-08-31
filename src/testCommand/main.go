package main

import "fmt"

func main() {
	fmt.Println(1, 2)
}

func add(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum + 1
}
