package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	sumData := []int{2, 3, 4, 5, 6}
	go sum(sumData[:2], c)
	go sum(sumData[2:], c)

	fmt.Println("sum is : ", <-c)
	fmt.Println("sum is : ", <-c)
}

func sum(a []int, c chan int) {
	data := 0
	for _, v := range a {
		data += v
	}

	c <- data

}
