package main

import (
	"fmt"
	"time"
)

//ticker is used to execute Go code at some point in the future.
func main() {
	timer1 := time.NewTimer(time.Second)
	<-timer1.C

	fmt.Println("1st timer")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("2nd timer")
	}()
	//you can cancel the timer before it fires
	stop := timer2.Stop()
	if stop {
		fmt.Println("stop")
	}
	time.Sleep(time.Second)
}
