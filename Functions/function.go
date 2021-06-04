package main

import (
	"fmt"
	"sort"
)

//Functions are a block of code that we can use repeatedly
func abc(a int, b int) int {
	return a * b
}
func abccba(a, b int, c string) int {
	fmt.Println(c)
	return a + b
}

//multiple return values in function
func ab() (int, int) {
	return 6, 7
}

//variadic functions are those which take irresptive number of arguments of same time
func varfu(abc ...int) {
	fmt.Println(abc)
	sort.Ints(abc)
	fmt.Println(abc)
}
func main() {
	res := abc(11, 4)
	fmt.Println(res)
	fmt.Println(abccba(10, 4, "fun"))
	a, b := ab()
	fmt.Println(a, b)
	varfu(3, 1)
	varfu(6, 3, 9)
	abc := []int{10, 2, 54, 68, 34}
	varfu(abc...)
}
