package main

import "fmt"

func main() {
	x := make([]int, 10, 50)
	fmt.Println(x)

	x[4] = 44
	x[5] = 55
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 111)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 112)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
