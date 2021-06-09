package main

import (
	"crypto/sha1"
	"fmt"
)

//secure hash algorithm
func main() {
	s := "Golang was first created by Google"
	g := sha1.New()
	g.Write([]byte(s))
	// The argument to Sum can be used to append to an existing byte slice
	bs := g.Sum(nil)
	fmt.Println(s)
	fmt.Println(bs)
}
