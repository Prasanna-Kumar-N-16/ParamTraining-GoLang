package main

import (
	"fmt"
	"time"
)

//which we’ll run several concurrent instances
func workers(id int, job <-chan int, res chan<- int) {
	for j := range job {
		fmt.Printf("%d worker started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("%d worker finished job %d\n", id, j)
		res <- j + 1
	}

}
func main() {
	numjobs := 5
	job := make(chan int, numjobs)
	res := make(chan int, numjobs)

	for j := 1; j <= 5; j++ {
		job <- j
	}
	for i := 1; i <= 3; i++ {
		go workers(i, job, res)
	}
	for k := 0; k < 5; k++ {
		<-res
	}

}
