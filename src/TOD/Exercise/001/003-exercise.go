package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)

	fmt.Println(s)
}
