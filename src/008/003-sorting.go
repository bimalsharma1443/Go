package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println("before sorting : ", xi)
	sort.Ints(xi)
	fmt.Println("after sorting : ", xi)

	fmt.Println("before sorting : ", xs)
	sort.Strings(xs)
	fmt.Println("after sorting : ", xs)

}
