package main

import (
	"fmt"
	"os"
	"strings"
)

//An environment variable is a dynamic-named value that
//can affect the way running processes will behave on a computer.
//Environment variables are a universal mechanism for
//conveying configuration information to Unix programs.
func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
