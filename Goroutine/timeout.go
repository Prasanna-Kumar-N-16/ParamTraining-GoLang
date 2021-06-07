package main

//timeouts are those used when we want to bound to specific execution time
import (
	"fmt"
	"time"
)

//the channel is buffered, so the send in the goroutine is nonblocking.
//This is a common pattern to prevent goroutine leaks in case the channel is never read.
func main() {
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 2)
	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "Param"
	}()
	select {
	case m1 := <-chan1:
		fmt.Println(m1)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 1")
	}
	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "Vaiydee"
	}()
	select {
	case m1 := <-chan2:
		fmt.Println(m1)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 2")
	}
}
