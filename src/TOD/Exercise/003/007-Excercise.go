package main

import "fmt"

func main() {

	i := 10

	if i < 0 {
		fmt.Println("value less than zero")
	} else if i <= 10 && i >= 0 {
		fmt.Println("value between zero and ten")
	} else {
		fmt.Println("value is : ", i)
	}

}
