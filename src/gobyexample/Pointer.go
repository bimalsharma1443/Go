package main

import "fmt"

func pointerIncrem(i int) {
	i++;
}

func valueIncrement(i *int) {
	*i++;
}

func main() {

	i := 1
	fmt.Println(i)

	pointerIncrem(i)
	fmt.Println(i)

	valueIncrement(&i)
	fmt.Println(i)
}