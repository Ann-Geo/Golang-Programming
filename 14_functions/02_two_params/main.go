package main

import "fmt"

func main() {
	greet("Jane", "Jim")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
