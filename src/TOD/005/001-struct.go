package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{
		firstName: "bimal",
		lastName:  "sharma",
		age:       29,
	}

	p2 := person{
		firstName: "sudeep",
		lastName:  "Prajapati",
		age:       28,
	}

	x := []person{p1, p2}
	fmt.Println(x)
}
