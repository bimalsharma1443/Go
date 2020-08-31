package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	encodeStd := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodeStd)

	decodeStd := base64.StdEncoding.DecodeString(encodeStd)
	fmt.Println(decodeStd)

}
