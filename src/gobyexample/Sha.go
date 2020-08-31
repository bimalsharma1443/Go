package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	stringData := "this is sha string"

	sha := sha1.New()

	sha.Write([]byte(stringData))

	bs := sha.Sum(nil)

	fmt.Println(stringData)
	fmt.Printf("%x\n", bs)

}
