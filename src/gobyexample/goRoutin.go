package main

import (
	"fmt"
	"time"
)

func f(msg string) {
	fmt.Println(msg)
}

func main() {

	f("this is a data")

	go func(msg string) {
		fmt.Println(msg)
	}("this is a message")

	time.Sleep(time.Second)
	fmt.Println("done")

}
