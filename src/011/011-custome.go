package main

import (
	"fmt"
	"log"
)

type user struct {
	name string
	err  error
}

func (u user) Error() string {
	return fmt.Sprintf("value is not corect for user %v and error is %v", u.name, u.err)
}

// type norgateMathError struct {
// 	lat  string
// 	long string
// 	err  error
// }

// func (n norgateMathError) Error() string {
// 	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
// }

func main() {

	_, e := sqrt(-10.23)

	if e != nil {
		log.Println(e)
	}

}

func sqrt(f float64) (float64, error) {

	if f < 0 {
		// nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		// return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
		name := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, user{"Bimal", name}

	}
	return 4, nil
}
