package main

import "fmt"

type bimal int

var x bimal

func main() {
	fmt.Println("value is : ", x)
	fmt.Printf("%t\n", x)
	x = 42
	fmt.Println("value is : ", x)

}
