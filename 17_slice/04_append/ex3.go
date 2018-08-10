package main

import "fmt"

func main() {
	first := []string{"Monday", "Tuesday"}
	second := []string{"Wednesday", "Thursday", "Friday"}

	first = append(first, second...)
	fmt.Println(first)
}
