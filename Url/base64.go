package main

import (
	B64 "encoding/base64"
	"fmt"
)

func main() {
	s := "This is Golang B64"
	h := B64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(h)
	i, _ := B64.StdEncoding.DecodeString(h)
	fmt.Println(i)
	j := B64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println(j)
	k, _ := B64.URLEncoding.DecodeString(j)
	//Decoding may return an error, which you can check if you don’t already know the input to be well-formed.
	fmt.Println(k)
}
