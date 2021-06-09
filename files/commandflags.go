package main

import (
	"flag"
	"fmt"
)

//Go provides a flag package supporting basic command-line flag parsing
func main() {

	word := flag.String("word", "foo", "a string")
	numb := flag.Int("numb", 42, "an int")
	bool := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *word)
	fmt.Println("numb:", *numb)
	fmt.Println("fork:", *bool)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
