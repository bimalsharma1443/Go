package main

import (
	"fmt"
	"math"
)

type circle struct {
	Radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.Radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	c1 := circle{Radius: 5}
	fmt.Println(c1)
	info(c1)
	fmt.Println(c1.area())

}
