package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.ParseInt("340040", 0, 64)
	fmt.Println(i)

	f, _ := strconv.ParseFloat("1.23", 64)
	fmt.Println(f)

	ui, _ := strconv.ParseUint("-123", 0, 64)
	fmt.Println(ui)
}
