package main

import "fmt"

func main() {
	defer boo()
	fmt.Println(soo())

}

func boo() string {
	fmt.Println("calling boo")
	return "boo"
}

func soo() string {
	fmt.Println("calling soo")
	return "soo"
}
