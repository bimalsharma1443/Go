package main

import "fmt"

func main() {
	// create map
	// map[key-type]value-type
	m := make(map[string]int)

	fmt.Println(" map is : ", m)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println(" map is : ", m)
	fmt.Println(" length of map ", len(m))

	v := m["k1"]
	fmt.Println(" value of k1 is", v)

	delete(m, "k2")
	fmt.Println(" map is : ", m)

	n := map[string]int{"abc": 1, "xyz": 2}
	fmt.Println(" map is : ", n)

}
