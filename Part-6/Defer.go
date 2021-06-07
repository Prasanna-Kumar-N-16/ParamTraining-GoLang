package main

import "fmt"

func add(i, j int) {
	fmt.Println(i + j)
}
func printstring() {
	fmt.Println("Hi")
}
func main() {
	//defer statements delay the execution of the function or
	// method or an anonymous method until the nearby functions returns
	add(21, 22)
	defer add(21, 22)
	printstring()
}
