package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"d", "e", "a"}
	fmt.Println("String : ", strs)
	sort.Strings(strs)
	fmt.Println("String : ", strs)

	ints := []int{7, 2, 4}
	fmt.Println("ints : ", ints)
	sort.Ints(ints)
	fmt.Println("ints : ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println(s)

	sort.IntSlice(ints)
	fmt.Println(ints)

}
