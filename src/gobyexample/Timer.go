package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("code is start working")

	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("time is fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("time 2 is fired")
	}()

	stop := timer2.Stop()

	if stop {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

}
