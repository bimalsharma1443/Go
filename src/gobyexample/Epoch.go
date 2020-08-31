package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	sec := now.Unix()
	bin := now.UnixNano()

	fmt.Println(sec)
	fmt.Println(bin)

}
