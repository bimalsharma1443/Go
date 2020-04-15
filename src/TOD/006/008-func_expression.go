package main

import "fmt"

func main() {

	fmt.Println("checking func expression")

	f := func() {
		fmt.Println("changes fun ran")
	}
	f()

	c := func(x int) {
		fmt.Println("x value is : ", x)
	}
	c(42)

}
