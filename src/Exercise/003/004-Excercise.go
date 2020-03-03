package main

import "fmt"

func main() {

	i := 1980

	for {

		fmt.Println(i)
		if i == 1990 {
			break
		}

		i++
	}
}
