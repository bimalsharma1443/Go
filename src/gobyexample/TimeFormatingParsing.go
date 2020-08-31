package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now()
	p(now.Format(time.RFC3339))

	p(now.Format("3:04 PM"))

	p(now.Format("2006-01-02"))
	p(now.Format("2019-12-12 3:04 PM"))
}
