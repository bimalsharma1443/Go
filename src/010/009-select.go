package main

import "fmt"

func main() {

	e := make(chan int)
	o := make(chan int)
	q := make(chan int)

	go send(e, o, q)

	recieve(e, o, q)

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	close(e)
	close(o)
	q <- 1

}

func recieve(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println(v)
		case v := <-o:
			fmt.Println(v)
		case <-q:
			return
		}

	}
}
