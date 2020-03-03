package main

import (
	"fmt"
	"log"
)

func main() {

	_, errs := sqrt(-10)
	if errs != nil {
		log.Fatalln(errs)
	}

}

func sqrt(i float64) (float64, error) {
	if i < 0 {
		err := fmt.Errorf("norgate math: square root of negative number : %v", i)
		return 0, err
	}
	return 100, nil
}
