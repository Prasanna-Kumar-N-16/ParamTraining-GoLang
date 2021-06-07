package main

import (
	"fmt"
	"time"
)

//tickers are for when you want to do something repeatedly at regular intervals
func main() {
	ticker := time.NewTicker(time.Second)
	boo := make(chan bool)
	go func() {
		for {
			select {
			case <-boo:
				return
			case t := <-ticker.C:
				fmt.Println(t)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	ticker.Stop()
	boo <- true
	fmt.Println("stop")
}
