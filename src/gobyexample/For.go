package main

import "fmt"

func main() {

	i := 1

	for i <  3 {
		fmt.Println(" i < 3",i)
		i++
	}

	for i = 1;i < 3; i++ {
		fmt.Println("the value is : ",i)
		i++
	}

	for i = 1; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i);
		}
	}

}