package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	//http.Get is a convenient shortcut around creating an http.Client object and calling its Get method
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//Print the HTTP response status
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
