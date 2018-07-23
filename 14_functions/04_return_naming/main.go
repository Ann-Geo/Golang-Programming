package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Jim"))
}

func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
