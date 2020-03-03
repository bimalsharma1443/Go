package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type bank struct {
	person
	bnk bool
}

func main() {

	sa1 := bank{
		person: person{
			"bimal",
			"sharma",
			27,
		},
		true,
	}

	fmt.Println(sa1)

}

// a SLICE allows you to group together VALUES of the same TYPE
