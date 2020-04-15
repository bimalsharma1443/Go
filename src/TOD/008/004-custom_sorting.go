package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type byAge []person
type byPeopel []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (p byPeopel) Len() int           { return len(p) }
func (p byPeopel) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byPeopel) Less(i, j int) bool { return p[i].Name < p[j].Name }

func main() {

	p1 := person{
		Name: "Bimal",
		Age:  29,
	}

	p2 := person{
		Name: "Sudip",
		Age:  27,
	}

	people := []person{p1, p2}

	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
	sort.Sort(byPeopel(people))
	fmt.Println(people)

}
