package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Error message is here %v", ce.info)
}

func main() {
	c1 := customErr{info: "bimal sharma "}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo run : ", e)

	fmt.Println("assertion: ", e.(customErr).info)
}
