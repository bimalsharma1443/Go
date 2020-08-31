package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	value  := time.Now().Unix()
	
	data := strconv.FormatInt(value,10)

	fmt.Printf("valie is %v and %t is \n",value,value)
	fmt.Printf("valie is %v and %t is ",data,data)
}