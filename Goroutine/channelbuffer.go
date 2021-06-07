package main

import "fmt"

//without a reciever and a limited count channel bufferring can be done
func main() {
	slack := make(chan string, 3)
	slack <- "beta"
	slack <- "general"
	slack <- "to contact"
	fmt.Println(<-slack)
	fmt.Println(<-slack)
	fmt.Println(<-slack)
}
