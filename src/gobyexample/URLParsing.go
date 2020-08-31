package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)

	if err != nil {
		panic(err)
	}

	fmt.Println(u)

	fmt.Println(u.Hostname())
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	fmt.Println(u.RawQuery)
	fmt.Println(url.ParseQuery(u.RawQuery))
}
