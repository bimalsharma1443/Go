package main

import "fmt"

func main() {

	fmt.Println("Please enter a number")
	var i int
	_, err := fmt.Scanln(&i)

	if err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case i < 0:
		fmt.Println("i is lower than zero")
	case i >= 0 && i <= 20:
		fmt.Println("i is between 0 and 20")
	case i > 20 && i <= 40:
		fmt.Println("i is between 20 and 40")
	case i > 40 && i <= 80:
		fmt.Println("i is between 40 and 80")
	case i > 80 && i <= 100:
		fmt.Println("i is between 80 and 100")
	default:
		fmt.Println("number is ", i)
	}

	whatIm := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm int")
		default:
			fmt.Printf("Dont know %T\n", t)
		}
	}

	whatIm(true)
	whatIm(1)
	whatIm("true is true")

}
