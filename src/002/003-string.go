package main

import "fmt"

func main() {

	s := "this is a string"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)

	for i := 0; i < len(s); i++ {
		fmt.Println("%x ", s[i])
	}

	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Println("%d ", s[i])
	}

	fmt.Println("")
	for i, v := range s {
		fmt.Printf("%#U \t %d", v, i)
	}

}
