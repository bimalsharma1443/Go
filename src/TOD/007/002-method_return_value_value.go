package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type calculation interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func check(c calculation) {
	fmt.Println(c.area())
}

func main() {

	c1 := circle{
		radius: 5,
	}

	check(c1)

}
