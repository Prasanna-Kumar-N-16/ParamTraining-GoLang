package main

import (
	"fmt"
	"math"
)

//Interfaces are named collections of method signatures.
//topval is the interface which connect two struct function types.
type topval interface {
	area() float64
	perim() float64
}
type rect struct {
	height float64
	width  float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func newtop(t topval) {
	fmt.Println(t)
	fmt.Println(t.area())
	fmt.Println(t.perim())
}
func main() {
	r := rect{height: 10, width: 20}
	c := circle{radius: 5}
	newtop(r)
	newtop(c)
}
