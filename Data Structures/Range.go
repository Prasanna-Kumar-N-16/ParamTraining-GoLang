package main

import (
	"fmt"
)

//Range  iterates over elements in a variety of data structures,with incexs and values.
func main() {
	s := []int{22, 33, 44, 55, 66}
	res := 0
	//Range used for arrays
	for i, num := range s {
		if i == 3 {
			fmt.Println(i)
		} else {
			res = res + num
		}
	}
	fmt.Println(res)
	//Range when used in maps
	abc := map[string]int{"abc": 1, "def": 2, "ghi": 3}
	for x, y := range abc {
		fmt.Printf("%s : %d\n", x, y)
	}
}
