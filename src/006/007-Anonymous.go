package main

import "fmt"

func main() {

	func() {
		fmt.Println("changes fun ran")
	}()

	func(x int) {
		fmt.Println("x value is : ", x)
	}(42)

}
