package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, error := os.Open("demo.txt")

	if error != nil {
		fmt.Println(error)
		return
	}

	defer f.Close()

	bs, error := ioutil.ReadAll(f)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(string(bs))

}
