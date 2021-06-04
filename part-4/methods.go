package main

import "fmt"

//Methods are types of struct functions
type comp struct {
	empname string
	empid   int
}

//concatinate function gets struct name and prints using . operation
func (c comp) concatinate() {
	fmt.Println(c.empname)
}
func (c comp) intretur() int {
	return c.empid * c.empid
}

func main() {
	c := comp{empname: "Prasanna", empid: 10}
	c.concatinate()
	fmt.Println(c.intretur())
	b := &c
	b.concatinate()
	fmt.Println(b.intretur())

}
