package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, error := os.Create("demo.txt")

	if error != nil {
		fmt.Println(error)
		return
	}

	defer f.Close()

	r := strings.NewReader("Bimal P Sharma")

	io.Copy(f, r)
}
