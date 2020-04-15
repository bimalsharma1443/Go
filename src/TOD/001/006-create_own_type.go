package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 40
	b = 43
	fmt.Println("the value is : ", a)
	fmt.Printf("%t\n", a)
	fmt.Println("the value is : ", b)
	fmt.Printf("%t\n", b)

	// uncomment give error becouse int and hotdog both are diffrent type.
	//  cannot use b (type hotdog) as type int in assignment
	/*a = b
	fmt.Println("the value is : ", a)
	fmt.Printf("%t\n", a)*/
}
