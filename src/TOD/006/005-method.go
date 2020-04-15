package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person
	etype string
}

func main() {
	fmt.Println("--------------start---------------")

	e1 := employee{
		person: person{
			"Bimal Sharma",
			30,
		},
		etype: "developer",
	}

	fmt.Println("used as value")
	e1.getDetail()

	re1 := &e1

	fmt.Println("used as reference")
	re1.getDetail()

	fmt.Println("--------------end---------------")

}

func (e employee) getDetail() {
	fmt.Println(e.person.name, " is a ", e.etype)
}
