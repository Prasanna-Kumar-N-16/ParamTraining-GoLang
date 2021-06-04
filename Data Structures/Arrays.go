package main

import "fmt"

//Arrays contain same data type
func main() {
	//1D Array
	var a [5]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	b := [5]int{23, 34, 45, 56, 67}
	fmt.Println(a)
	fmt.Println(b)
	//2D array
	var d [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = i + j
		}
	}
	fmt.Println(d)
}
