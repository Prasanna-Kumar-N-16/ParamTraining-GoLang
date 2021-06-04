package main

import (
	"errors"
	"fmt"
)

func func1(num int) (int, error) {
	if num == 36 {
		return -1, errors.New("bad choice")
	} else {
		return num + 3, nil
	}
}

func main() {
	arr := []int{22, 36, 42}
	for _, i := range arr {
		a, d := func1(i)
		if d != nil {
			fmt.Println("func1 didn't work", d)
		} else {
			fmt.Println("Func1 worked", a)
		}
	}
}
