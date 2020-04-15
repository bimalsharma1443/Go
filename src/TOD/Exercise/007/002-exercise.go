package main

import "fmt"

type person struct {
	first string
}

func main() {

	p1 := person{
		first: "vimal",
	}

	fmt.Println(p1)
	changeName(p1)
	fmt.Println(p1)
	changeNamePr(&p1)
	fmt.Println(p1)
}

func changeName(p person) {
	p.first = "bimal"
}

func changeNamePr(p *person) {
	(*p).first = "bimal"
}
