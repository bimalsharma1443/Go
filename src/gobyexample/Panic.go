package main

import "fmt"

func main() {
	var value int

	fmt.Println("Please eneter a number")
	_, err := fmt.Scanln(&value)

	if err != nil {
		panic("please enter a number")
		return
	}

	println("You have enter a numbe is : ", value)

}
