package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	permi() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return (r.width * r.height)
}

func (r rect) permi() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) permi() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area is : ", g.area())
	fmt.Println("permi is : ", g.permi())
}

func main() {

	fmt.Println("Please select option to calculate following")
	fmt.Println("1 for circle")
	fmt.Println("2 for rectangle")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		var radius float64
		fmt.Println("Please enter radius")
		fmt.Scanln(&radius)

		c := circle{
			radius,
		}
		measure(c)

	case 2:
		var width float64
		var height float64
		fmt.Println("Please enter width")
		fmt.Scanln(&width)
		fmt.Println("Please enter height")
		fmt.Scanln(&height)

		r := rect{
			width, height,
		}
		measure(r)

	}
}
