package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, " event number")
		}
	}
	check()
}

func check() {
	fmt.Println(" in side a check")
}
