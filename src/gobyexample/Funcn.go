package main

import "fmt"

// or fun plus(a,b int) (int)
func plus(a int, b int) int {
	return a + b
}

func add(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {

	sum := plus(3, 5)
	fmt.Println("Sum is : ", sum)

	sum = add(1, 2, 3, 5)
	fmt.Println("Sum is : ", sum)

	value := []int{4, 5, 6, 7, 8}
	sum = add(value...)
	fmt.Println("Sum is : ", sum)

}
