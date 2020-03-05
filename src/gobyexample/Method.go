package main

import "fmt"

type rect struct {
	length , breadth int
}

func (r *rect) area() (int) {
	return r.length*r.breadth 
} 

func (r *rect) per() (int) {
	return 2*(r.length+r.breadth)
}



func main() {

	l := 0
	b := 0

	fmt.Println("Please enter length : ")
	fmt.Scanln(&l)

	fmt.Println("Please enter breadth : ")
	fmt.Scanln(&b)

	rectData := rect {
		l,b,
	}

	area := rectData.area()
	permi := rectData.per()

	fmt.Println("area is ", area)
	fmt.Println("perim is ",permi)

}