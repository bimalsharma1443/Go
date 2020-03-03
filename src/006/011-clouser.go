package main

import "fmt"

func main() {

	a := increment()
	b := increment()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
}

func increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}
