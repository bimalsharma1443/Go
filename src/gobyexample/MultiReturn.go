package main

import "fmt"

func sumAndMulti(value ...int) (int, int) {
	sum := 0
	multi := 1

	for _, v := range value {
		sum += v
		multi *= v
	}

	return sum, multi

}

func main() {

	sum, multi := sumAndMulti(1, 2, 3, 4)

	fmt.Printf("Sum is %v and Multi is %v", sum, multi)

}
