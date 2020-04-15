package main

import (
	"fmt"
)

func send(max int, message chan<- int) {
	for count := 0; count < max; count++ {
		message <- count * 2
	}
	close(message)
}

func recieve(message <-chan int) {
	for value := range message {
		fmt.Println(value)
	}
}

func main() {

	i := 10

	message := make(chan int)

	go send(i, message)

	recieve(message)

}
