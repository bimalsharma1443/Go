package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now()

	p(now)

	then := time.Date(2020,4,24,14,23,00,651387237,time.UTC)

	p(then.Year())


}