package main

import "fmt"

func main() {

	name := []string{"MH", "GJ", "GA", "MP", "RJ"}
	fmt.Println("Size of slice", len(name))
	fmt.Println("Capacity of slice", cap(name))

	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

}
