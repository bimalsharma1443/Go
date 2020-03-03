package main

import "fmt"

func main() {

	d := struct {
		name     string
		friend   map[int]string
		favDrink []string
	}{
		name: "Bimal",
		friend: map[int]string{
			1: "Sudeep",
			2: "Sandeep",
		},
		favDrink: []string{
			"water",
			"juice",
		},
	}

	fmt.Println(d)

	for i, v := range d.friend {
		fmt.Println(i, v)
	}

	for i, v := range d.favDrink {
		fmt.Println(i, v)
	}

}
