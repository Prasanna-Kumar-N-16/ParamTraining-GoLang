package main

//using a blocking receive we can wait  for a goroutine to finish
import (
	"fmt"
	"time"
)

/*func nfunc(nechan chan string) {
	fmt.Println("1st run")
	time.Sleep(10 * time.Second)
	fmt.Println("2nd run")
	nechan <- "Param"
}*/
func main() {
	newchan := make(chan string, 1)
	//go nfunc(newchan)
	go func(nechan chan string) {
		fmt.Println("1st run")
		time.Sleep(10 * time.Second)
		fmt.Println("2nd run")
		nechan <- "Param"
	}(newchan)
	fmt.Println(<-newchan)
}
