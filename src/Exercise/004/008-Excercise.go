package main

import "fmt"

func main() {

	x := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range x {
		fmt.Println("the value of this record : ", k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

}
