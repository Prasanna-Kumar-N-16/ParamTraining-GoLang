package main

import (
	"fmt"
)

func main() {
	format := make(chan int, 1)
	go func() {
		//Send a value into a channel using the channel <- value syntax.
		format <- 420
	}()
	//The <-channel syntax receives a value from the channel.
	//robb := <-format
	fmt.Println(<-format)
	rangee := make(chan int, 2)
	rangee <- 1
	rangee <- 2
	close(rangee)
	for elem := range rangee {
		fmt.Println(elem)

	}
}
