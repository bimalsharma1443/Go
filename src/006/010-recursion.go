package main

import "fmt"

func main() {

	v := recu(4)
	fmt.Println(v)
}

func recu(v int) int {
	if v == 1 {
		return 1
	}

	return v * recu(v-1)
}
