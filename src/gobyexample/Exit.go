package main

import "fmt"
import "os"

func main() {
	defer fmt.Println("!")

	os.Exit(100000000000)
}
