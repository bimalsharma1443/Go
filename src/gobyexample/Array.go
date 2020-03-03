package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("Array is : ",a)

	a[4] = 20
	fmt.Println("Array is : ",a)

	fmt.Println("length of array : ",len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array is : ",b)

	var towD [5][5]int
	for i := 0; i < 5 ;i++ {
		for j := 0; j < 5 ;j++ {
			towD[i][j] = i + j
		}
	}

	fmt.Println(towD)


}