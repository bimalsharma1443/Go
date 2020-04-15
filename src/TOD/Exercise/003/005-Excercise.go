package main

import "fmt"

func main() {

	i := 10

	for i <= 100 {

		if i%4 == 0 {
			fmt.Println(i)
		}

		i++
	}

}
