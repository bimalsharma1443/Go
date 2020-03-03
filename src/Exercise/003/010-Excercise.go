package main

import "fmt"

func main() {

	data := 'b'

	switch data {
	case 'a':
		fmt.Println("a is printing")
	case 'b':
		fmt.Println("b is printing")
	default:
		fmt.Printf("%v is printing", data)
	}

}
