package main

import (
	"fmt"
	"sort"
)

//sort is something which we increase / decrease accordingly
type size []string

func (s size) Len() int {
	return len(s)
}
func (s size) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s size) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
func main() {
	/*strings := []string{"prasanna", "meghana", "param"}
	num := []int{23, 12, 4}
	sort.Strings(strings)
	sort.Ints(num)
	fmt.Println(strings)
	fmt.Println(num)*/
	str := []string{"prasanna", "meghana", "param"}
	sort.Sort(size(str))
	fmt.Println(str)
}
