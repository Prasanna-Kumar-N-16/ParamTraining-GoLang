package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//we often want to create data that isn’t needed after the program exits.
	//Temporary files and directories are useful for this purpose since they don’t pollute the file system over time.
	t, err := ioutil.TempFile("", "Go")
	check(err)
	fmt.Println("Temp file name:", t.Name())
	defer os.Remove(t.Name())
	dirname, err := ioutil.TempDir("", "GOtempdir")
	check(err)
	fmt.Println(dirname)
	defer os.RemoveAll(dirname)
}
