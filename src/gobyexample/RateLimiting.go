package main

import (
	"fmt"
	"time"
)

func main() {
	request := make(chan int, 5)

	for i := 0; i < 5; i++ {
		request <- 1
	}
	close(request)

	limiter := time.Tick(500 * time.Millisecond)

	for req := range request {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

}
