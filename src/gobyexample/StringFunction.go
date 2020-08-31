package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	stringData := "This is a testing a string"

	stringArray := []string{"String 1", "String 2", "String 3"}

	p("Contains:  ", s.Contains(stringData, "ing"))
	p("Count:     ", s.Count(stringData, "g"))
	p("HasPrefix: ", s.HasPrefix(stringData, "t"))
	p("HasSuffix: ", s.HasSuffix(stringData, "g"))
	p("Index:     ", s.Index(stringData, "ing"))
	p("Join:      ", s.Join(stringArray, ","))
	p("Repeat:    ", s.Repeat(stringData+". ", 3))
	p("Replace:   ", s.Replace(stringData, "t", "tt", -1))
	p("Replace:   ", s.Replace(stringData, "t", "tt", 1))
	p("Split:     ", s.Split(stringData, " "))
	p("ToLower:   ", s.ToLower(stringData))
	p("ToUpper:   ", s.ToUpper(stringData))
	p()

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

}
