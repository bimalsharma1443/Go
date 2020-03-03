package main2

import "fmt"

func main() {

	result, err := sum(1, 2, 3)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}

func sum(values ...int) (int, error) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum, nil
}
