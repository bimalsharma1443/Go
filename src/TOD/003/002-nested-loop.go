package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf(" outer loop : %v\n", i)
		for j := 0; j < 10; j++ {
			fmt.Printf(" inner loop : %v\t\t\n", j)
		}
	}

}
