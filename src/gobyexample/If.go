package main

import "fmt"

func main() {

	fmt.Println("Please enter a number")

	var i int

	_, err := fmt.Scanln(&i)

	if err != nil {
		fmt.Println(err)
		return
	}

	if i%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if i%4 == 0 {
		fmt.Println("number is divisiable by 4")
	}

	if i < 0 {
		fmt.Println("number is -ve")
	} else if i > 0 {
		fmt.Println("number is +ve")
	} else {
		fmt.Println("number is zero")
	}
}
