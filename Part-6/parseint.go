package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("9.653", 64)
	fmt.Println(f)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
}
