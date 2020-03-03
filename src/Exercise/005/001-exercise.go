package main

import "fmt"

type person struct {
	first string
	last  string
	ice   []string
}

func main() {

	p1 := person{
		first: "bimal",
		last:  "sharma",
		ice: []string{
			"Mango",
			"vanilla",
		},
	}

	p2 := person{
		first: "bimal",
		last:  "sharma",
		ice: []string{
			"strawberry",
			"Mango",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)

	for i, v := range p1.ice {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)

	for i, v := range p2.ice {
		fmt.Println(i, v)
	}

}
