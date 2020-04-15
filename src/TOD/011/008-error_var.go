package main

import (
	"errors"
	"log"
)

var err = errors.New("norgate math: square root of negative number")

func main() {

	_, errs := sqrt(-10)
	if errs != nil {
		log.Fatalln(errs)
	}

}

func sqrt(i float64) (float64, error) {
	if i < 0 {
		return 0, err
	}
	return 100, nil
}
