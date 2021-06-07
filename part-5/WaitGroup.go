package main

//waitgroups blocks until any goroutines within that WaitGroup have successfully executed.
import (
	"fmt"
	"sync"
	"time"
)

func funct(i int, a *sync.WaitGroup) {
	//we call .Done() within any goroutine to signal the end of its' execution.
	defer a.Done()
	fmt.Printf("%d\n", i)
	time.Sleep(time.Second)
	fmt.Printf(" next %d\n", i)
}

func main() {
	var a sync.WaitGroup
	for i := 1; i <= 5; i++ {
		//We  call .Add(1) on our WaitGroup to set the number of goroutines we want to wait for
		a.Add(1)
		go funct(i, &a)
	}
	a.Wait()
}
