package main

import (
	"fmt"
	"time"
)

//select lets you wait on multiple channel operations.
func main() {
	chanone := make(chan string, 2)
	chantwo := make(chan int, 2)
	go func(chanone chan string) {
		time.Sleep(2 * time.Second)
		chanone <- "abcd"
	}(chanone)
	go func(chantwo chan int) {
		time.Sleep(3 * time.Second)
		chantwo <- 2
	}(chantwo)
	for i := 0; i < 2; i++ {
		//We’ll use select to wait both of these values simultaneously, printing each one as it arrives
		select {
		case m1 := <-chanone:
			fmt.Println(m1)
		case m2 := <-chantwo:
			fmt.Println(m2)
		}
	}
}
