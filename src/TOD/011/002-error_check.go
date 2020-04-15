package main

import "fmt"

func main() {
	var answer1, answer2, answer3 string

	fmt.Print("Your First Name : ")
	_, error := fmt.Scanln(&answer1)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Print("Your Middle Name : ")
	_, error = fmt.Scanln(&answer2)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Print("Your Last Name : ")
	_, error = fmt.Scanln(&answer3)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(answer1, answer2, answer3)
}
