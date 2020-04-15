package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4}

	sum := 0
	for _, v := range num {
		sum += v
	}

	fmt.Println(sum)

	kvs := map[string]string{"a": "apple", "b": "mango"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key : ", k)
	}

	for k, v := range "golang" {
		fmt.Println(k, v)
	}

}
