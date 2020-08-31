package main

import (
	"encoding/json"
	"fmt"
)

type db struct {
	value  int
	fruits []string
}

func main() {

	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(intB)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(2.34)
	fmt.Println(floatB)
	fmt.Println(string(floatB))

	strB, _ := json.Marshal("this is a string")
	fmt.Println(strB)
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(slcB)
	fmt.Println(string(slcB))

	var data []string
	json.Unmarshal(slcB, &data)
	fmt.Println(data)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

}
