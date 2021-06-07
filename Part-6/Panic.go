package main

import "os"

//A panic typically means something went unexpectedly wrong.
func main() {
	panic("Alert")
	_, err := os.Create("/temp/file")
	if err != nil {
		panic(err)
	}

}
