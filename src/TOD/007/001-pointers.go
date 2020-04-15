package main

import "fmt"

func main() {
	a := 4
	fmt.Println(a)
	increment(a)
	fmt.Println(a)
	incrementPointrt(&a)
	fmt.Println(&a)
	fmt.Println(a)

	b := &a
	fmt.Println(b)
	*b++
	fmt.Println(*b)
}

func increment(a int) {
	a++
}

func incrementPointrt(a *int) {
	*a++
}
