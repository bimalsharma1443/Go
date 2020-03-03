package main

import "fmt"

func main() {

	m := map[string]int{
		"Bimal":  29,
		"Sudeep": 28,
	}
	fmt.Println(m)
	fmt.Println(m["Bimal"])
	fmt.Println(m["Sudeep"])

	m["Sahel"] = 27

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "Bimal")

	fmt.Println(m["Bimal"])

	for k, v := range m {
		fmt.Println(k, v)
	}

}
