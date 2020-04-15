package main

import (
	"fmt"
)

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum(nums...)

}

func sum(nums ...int) {
	sum := 0
	for i, v := range nums {
		fmt.Println(i, v)
		sum += v
	}
	fmt.Println("Sum of ", nums, " number is : ", sum)
}

// a SLICE allows you to group together VALUES of the same TYPE
