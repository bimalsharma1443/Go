package main

import (
	"errors"
	"fmt"
)

func call(arg int) (int, error) {
	if arg == 24 {
		return -1, errors.New("can't work with 24")
	} else {
		return arg + 3, nil
	}
}

func main() {
	i := 28

	result, err := call(i)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
