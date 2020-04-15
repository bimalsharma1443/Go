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
		last:  "sharma1",
		ice: []string{
			"strawberry",
			"Mango",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for i, v := range m {
		fmt.Println(i)
		fmt.Println(v)
		fmt.Println("----------------------------")
	}

}
