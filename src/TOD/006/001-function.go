package main

import "fmt"

func main() {
	resp := plus(3, 4)
	fmt.Println(resp)

	resp = plusPlus(3, 4, 5)
	fmt.Println(resp)

	sum, sub := plusMinus(3, 4)
	fmt.Println(sum)
	fmt.Println(sub)
}

// function
func plus(a int, b int) int {
	return a + b
}

// no need to declare multiple type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func plusMinus(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}
