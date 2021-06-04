package main

import "fmt"

//slices provide powerful interface in Go when compared to arrays
func main() {
	d := make([]int, 4)
	for i := 0; i < len(d); i++ {
		d[i] = i * i
	}
	fmt.Println(d)
	//append a special feature in slices
	d = append(d, 16)
	d = append(d, 25, 36)
	fmt.Println(d)
	t := make([]int, len(d))
	//copy one slice to other slice
	copy(t, d)
	fmt.Println(t)
	fmt.Println(d[:3])
	fmt.Println(t[2:])
	//two d slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		//length of inner slice can vary irrespective of of array
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
