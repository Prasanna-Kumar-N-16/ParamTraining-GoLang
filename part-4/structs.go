package main

import "fmt"

// struct  is collections of fields.It is useful for grouping data together to form records.
type comp struct {
	empname string
	id      int
}

//a function link with struct
func newcomp(newemp string, newid int) *comp {
	p := comp{empname: newemp, id: newid}
	return &p
}

func main() {
	fmt.Println(comp{empname: "xyz", id: 105})
	fmt.Println(comp{"param", 000})
	fmt.Println(comp{empname: "vaidee"})
	fmt.Println(newcomp("Anand Rathi", 1))
	acc := comp{empname: "Jitendra", id: 1}
	//structs can be accessed using dots
	fmt.Println(acc.empname, acc.id)
}
