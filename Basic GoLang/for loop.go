package main

import "fmt"

func main() {
	//Used for loop with if , if else , break and continue statements
	for i := 0; i <= 10; i++ {
		if i == 8 {
			break
		} else if i == 3 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
