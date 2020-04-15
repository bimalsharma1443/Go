package main

import "fmt"

func main() {

	s := make([]int, 2)
	fmt.Println(s)

	s[0] = 12
	fmt.Println(s)

	s[1] = 13
	fmt.Println(s)

	s = append(s, 14, 15)
	fmt.Println(s)

	d := []int{6, 17, 18}
	fmt.Println(d)

	// if you use like this append(s,d) it gives an error.
	s = append(s, d...)
	fmt.Println(s)

	c := make([]int, 3)
	fmt.Println("before copy : ", c)
	copy(c, s)
	fmt.Println("after copy : ", c)

	c1 := make([]int, 10)
	fmt.Println("before copy : ", c1)
	copy(c1, s)
	fmt.Println("after copy : ", c1)

	fmt.Println("main array : ", c1)
	fmt.Println("slice detail between 1 to 3 ", c1[1:3])
	fmt.Println("slice detail between 0 to 3 ", c1[0:3])
	fmt.Println("slice detail between start to 6 ", c1[:6])
	fmt.Println("slice detail between 6 to end ", c1[6:])

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j < (i + 1); j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("Slice 2D array : ", twoD)

}
