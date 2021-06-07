//We can specify if a channel is meant to only send or receive values when used as function prameters.
//This specificity increases the type-safety of the program.
package main

import "fmt"

func funcone(chanone chan<- int, num int) {
	chanone <- num
}
func functwo(chantwo chan<- int, chanone <-chan int) {
	temp := <-chanone
	chantwo <- temp
}
func main() {
	chanone := make(chan int, 5)
	chantwo := make(chan int, 5)
	for i := 0; i < 5; i++ {
		funcone(chanone, i+1)
		functwo(chantwo, chanone)
		fmt.Println(<-chantwo)
	}
}
