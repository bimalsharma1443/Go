package main

import "fmt"

func main() {

	c := make(chan int)

	go send(c)

	recieve(c)

}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func recieve(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}

}
