package main

import "fmt"

func main() {
	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "bimal",
		lastName:  "sharma",
		age:       29,
	}

	fmt.Println(p1)
}
