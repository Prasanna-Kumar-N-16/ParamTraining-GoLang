package main

//The standard library’s strings package provides many useful string-related functions.
import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p(s.Contains("abcd", "z"))
	p(s.Count("abcacb", "a"))
	p(s.HasPrefix("param", "pa"))
	p(s.HasSuffix("Rathi", "hi"))
	p(s.ToLower("PRASAnna"))
	p(s.Split("na/me", "/"))
	p(s.Repeat("e", 4))
}
