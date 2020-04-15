package main

import "fmt"

func main() {

	switch {
	case (2 == 2):
		fmt.Println("default true statement")
	}

	a := 20

	switch {
	case (a < 0):
		fmt.Println("value is less than : ")
	case (a > 0 && a <= 10):
		fmt.Println("value is between 0 and 10")
	case (a > 10):
		fmt.Println("value is greater than 10")
	case (a < 10 && a >= 20):
		fmt.Println("value is between 10 and 20")

	}
}
