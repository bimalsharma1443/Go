package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(len(x))

	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(x[5])

	for i, v := range x {
		fmt.Println(i, v)
	}

}
