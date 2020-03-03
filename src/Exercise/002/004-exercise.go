package main

import "fmt"

var a = 23

func main() {
	fmt.Println(a)

	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)

}
