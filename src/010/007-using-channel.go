package main

import "fmt"

func main() {

	c := make(chan int)

	go send(c)

	recieve(c)

}

func send(c chan<- int) {
	c <- 42
}

func recieve(c <-chan int) {
	fmt.Println(<-c)
}
