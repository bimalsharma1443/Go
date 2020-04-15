package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	fmt.Println("Message sending is start")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awating a signal")
	<-done
	fmt.Println("existing")

}
