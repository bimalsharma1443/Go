package main

import (
	"fmt"
	"strings"
)

func Index(t string, vc []string) int {
	for i, v := range vc {
		if v == t {
			return i
		}
	}
	return -1
}

func Filter(vc []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vc {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {

	str := []string{"Apple", "Orange", "Kiwee", "Mango", "Pineapple"}

	fmt.Println(Index("apple", str))

	fmt.Println(Filter(str, func(v string) bool {
		return strings.Contains(v, "e")
	}))

}
