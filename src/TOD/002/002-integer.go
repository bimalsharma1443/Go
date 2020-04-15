package main

import (
	"fmt"
	"runtime"
)

var x int

func main() {

	fmt.Println("the value is : ", x)
	x = 12
	fmt.Println("the value is : ", x)
	fmt.Printf("%T\n", x)
	y := 23.3456
	fmt.Println("the value is : ", y)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
