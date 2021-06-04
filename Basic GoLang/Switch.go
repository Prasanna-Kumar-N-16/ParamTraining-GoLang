package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		switch i {
		case 1:
			fmt.Println(i)
		case 2:
			fmt.Println(i)
		case 3:
			fmt.Println(i)
			continue
		default:
			fmt.Println("Last One")
		}
	}
}
