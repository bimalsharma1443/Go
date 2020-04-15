package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	permi() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	height, width float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) permi() float64 {
	return 2 * c.radius * math.Pi
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) permi() float64 {
	return 2 * (r.height + r.width)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.permi())
}

func main() {

	r := rectangle{
		height: 2,
		width:  3,
	}

	c := circle{
		radius: 3,
	}

	measure(r)
	measure(c)

}
