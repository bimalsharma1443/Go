package main

import "fmt"

func send(e, o, q chan<- int) {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}

	}

	q <- 1
}

func recieve(e, o, q <-chan int) {
	for {
		select {
		case msg1 := <-e:
			fmt.Println("even : ", msg1)
		case msg2 := <-o:
			fmt.Println("odd : ", msg2)
		case <-q:
			return

		}
	}

}

func main() {

	e := make(chan int)
	o := make(chan int)
	q := make(chan int)

	go send(e, o, q)

	recieve(e, o, q)

}
