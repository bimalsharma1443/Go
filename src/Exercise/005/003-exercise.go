package main

import "fmt"

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	v1 := truck{
		vehicle{
			2,
			"red",
		},
		true,
	}

	v2 := sedan{
		vehicle{
			4,
			"black",
		},
		true,
	}

	fmt.Println(v1.vehicle)
	fmt.Println(v1.fourWheel)
	fmt.Println(v2.vehicle)
	fmt.Println(v2.luxury)

}
