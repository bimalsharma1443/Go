package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	x = append(x, 6, 7, 8, 9)
	fmt.Println(x)

	y := []int{10, 11, 12, 13, 14}
	x = append(x, y...)
	fmt.Println(x)

	// delete from slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}
