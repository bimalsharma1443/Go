package main

import "fmt"

type bimal int

var x bimal
var y int

func main() {
	fmt.Println("value is : ", x)
	fmt.Printf("%t\n", x)
	x = 42
	fmt.Println("value is : ", x)

	y = int(y)
	fmt.Println("value is : ", y)
	fmt.Printf("%t\n", y)

}
