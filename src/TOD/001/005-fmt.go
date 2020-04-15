package main

import "fmt"

var x = 40

func main() {
	fmt.Println("the value is : ", x)
	fmt.Printf("%t\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)

	y := 123
	fmt.Printf("%#x\n", y)
	fmt.Printf("%t\t%b\n", y)

	z := fmt.Sprintf("%b\n", y)
	fmt.Println("the value is print: ", z)
}
