package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {

	s := `[{"Name":"Bimal","Age":32},{"Name":"Sudeep","Age":27}]`
	sb := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", sb)

	var people []person

	err := json.Unmarshal(sb, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

}
