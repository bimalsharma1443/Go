package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	fmt.Println("lets start crypto package.")

	genkey := make([]byte, 16)

	c, err := aes.NewCipher(genkey)

	if err != nil {
		panic(err)
	}

}
