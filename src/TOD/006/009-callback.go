package main

import "fmt"

func main() {

	i := []int{1, 2, 3, 4, 5}

	s := sum(i...)
	fmt.Println("sum of all number is : ", s)

	ev := even(sum, i...)
	fmt.Println("even sum is : ", ev)

	ov := odd(sum, i...)
	fmt.Println("odd sum is : ", ov)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func even(sum func(x ...int) int, data ...int) int {
	var even []int
	for _, v := range data {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return sum(even...)
}

func odd(sum func(x ...int) int, data ...int) int {
	var odd []int
	for _, v := range data {
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}
	return sum(odd...)
}
