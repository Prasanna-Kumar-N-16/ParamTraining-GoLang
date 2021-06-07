package main

import (
	"fmt"
	"time"
)

//Goroutines are light weight threads which executes simultaneosly and concurrently with other go routines
func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d\n", i)
	}
}
func main() {
	f("abc")
	go f("def")
	go func(s string) {
		fmt.Println(s)
	}("abc")
	time.Sleep(time.Minute)
	fmt.Println("done")
}
