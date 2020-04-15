package main

import (
	"fmt"
)

func main() {

	d, error := fmt.Println("Hi this is first error")

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(d)

}
