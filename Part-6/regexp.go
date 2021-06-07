package main

import (
	"fmt"
	"regexp"
)

var p = fmt.Println

func main() {
	//match, _ := regexp.MatchString("p([a-z]+)k", "peak")
	r, _ := regexp.Compile("p([a-z]+)k")
	//p(match)
	p(r.MatchString("peak"))
	//This finds the match for the regexp
	p(r.FindString("peak plunk"))
	p(r.FindStringIndex("peak"))
	p(r.FindStringSubmatch("peak plank"))
	// this will return information about the indexes of matches and submatches.
	p(r.FindStringSubmatchIndex("peak pork"))
	p(r.FindAllString("peak pork plank", -1))
	p(r.Match([]byte("peak")))
}
