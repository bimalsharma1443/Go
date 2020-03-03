package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, error := os.Create("log.log")

	if error != nil {
		fmt.Println(error)
	}

	defer f.Close()

	log.SetOutput(f)

	_, error = os.Open("checking_file.txt")

	if error != nil {
		// fmt.Println(error)
		log.Println(error)
		// log.Fatalln(error)
		// log.Panicln(error)
		// panic(error)
	}

	fmt.Println("end")

}
