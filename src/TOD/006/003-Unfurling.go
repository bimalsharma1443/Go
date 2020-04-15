package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3))

	d := []int{1, 2, 3, 4}

	fmt.Println(sum(d...))
}

func sum(num ...int) int {
	sum := 0
	for i, v := range num {
		fmt.Println(i, v)
		sum += v
	}
	return sum
}
