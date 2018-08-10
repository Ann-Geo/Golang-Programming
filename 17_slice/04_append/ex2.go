package main

import "fmt"

func main() {

	first := []int{1, 2, 3, 4, 5}
	second := []int{6, 7, 8, 9}

	first = append(first, second...)

	fmt.Println(first)
}
