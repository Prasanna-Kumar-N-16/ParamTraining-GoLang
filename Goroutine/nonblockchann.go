package main

import "fmt"

func main() {
	strings := make(chan string)
	integ := make(chan int, 1)
	select {
	case msg1 := <-strings:
		fmt.Println(msg1)
	default:
		fmt.Println("default 1")
	}
	integ <- 22
	select {
	case in := <-integ:
		fmt.Println(in)
	default:
		fmt.Println("default 2")
	}
	//channel has no buffer and there is no receiver so default got selected
	stri := "Go"
	select {
	case strings <- stri:
		fmt.Println(stri)
	default:
		fmt.Println("default 3")
	}
}
