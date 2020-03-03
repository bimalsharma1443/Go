package main

import "fmt"

func main() {

	s := "H"
	fmt.Printf("%s\n", s)

	b := []byte(s)
	fmt.Printf("%q\n", b)

	n := b[0]
	fmt.Println(n)

	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#x\n", n)
}
