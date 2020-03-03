package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	fmt.Println("Message sending is start")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	s := <-sigs

	fmt.Println("Got signal:",s)

}