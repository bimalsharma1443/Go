package main

import "fmt"

type Check struct {
	TypeData string
	status   bool
}

func main() {

	check := Check{"data", true}

	callData(check)
}

func callData(data interface{}) {
	fmt.Println(data)
}
