package main

import "fmt"

type person struct {
	name string
	age int
}

func newName(name string) *person {

	p := person{name : name}
	p.age = 42

	return &p;

}


func main() {

	p1 := person{
		name : "Bimal Sharma",
		age : 29,
	} 

	fmt.Println(p1)

	fmt.Println(person{"Satish",23})

	p2 :=  newName("Sudip")
	fmt.Println(p2)

}