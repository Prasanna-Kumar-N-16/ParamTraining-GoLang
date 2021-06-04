package main

import (
	"fmt"
)

//Maps are built-in associative data type sometimes called hashes or dicts in other languages
func main() {
	//Maps contain key value pair
	ma := make(map[string]int)
	for i := 0; i < 3; i++ {
		ma["A"] = i
	}
	ma["B"] = 23
	fmt.Println(len(ma))
	fmt.Println(ma)
	delete(ma, "A")
	fmt.Println(ma)
	//checks whether it is in map or not
	_, prs := ma["ras"]
	fmt.Println(prs)
	nma := map[string]int{"abc": 1, "xyz": 2}
	fmt.Println(nma)
}
