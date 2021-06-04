package main

import "fmt"

//Anonymous functions are those which  you use to define a function inline without having to name it.
func newclos() func() int {
	i := 0
	return func() int {
		i++
		i = i * i
		return i
	}
}
func main() {
	clos := newclos()
	fmt.Println(clos())
	fmt.Println(clos())
}
