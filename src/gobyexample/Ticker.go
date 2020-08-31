package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start time ticker")

	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Ticker at", t)
			}
		}
	}()

	time.Sleep(25 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("ticker Stop")

}
