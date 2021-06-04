package main

import "fmt"

//pointers refer to the address of variables
func noref(a int, b int) {
	a = 20
	b = 10
}

//pointers allow you to pass references to values and records within your program.
func ref(b *int, a *int) {
	*b = 20
	*a = 10

}
func main() {
	i := 10
	j := 20
	noref(i, j)
	fmt.Println(i, j)
	ref(&i, &j)
	fmt.Println(i, j)
	fmt.Println(&i, &j)

}
