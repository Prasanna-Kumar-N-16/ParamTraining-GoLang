package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	a := filepath.Join("c:/", "Users/", "radha/", "Desktop/", "GoLangg.txt")
	fmt.Println(a)
	fmt.Println(filepath.Join("c:/Users/radha/Desktop/", "Golangg.txt"))
}
