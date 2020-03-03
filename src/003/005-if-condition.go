package main

import "fmt"

func main() {
	if true {
		fmt.Println("this is a true value")
	}

	if false {
		fmt.Println("this is a false value")
	}

	if !true {
		fmt.Println("this is a true not value value")
	}

	if !false {
		fmt.Println("this is a false not value value")
	}

	if !(2 == 2) {
		fmt.Println("checking !2==2")
	}

	if 2 == 2 {
		fmt.Println("checking 2==2")
	}

	if 2 != 2 {
		fmt.Println("checking 2!=2")
	}

	if !(2 != 2) {
		fmt.Println("checking !(2!=2)")
	}

	x := 42

	if x == 42 {
		fmt.Println("value of x is match")
	} else {
		fmt.Println("value of x is not match")
	}

}
