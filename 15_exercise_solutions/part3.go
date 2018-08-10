package main

import "fmt"

func main() {
	greatest := findGreatest(3, 10, 2, 7, 0, 2)
	fmt.Println(greatest)
}

func findGreatest(val ...int) int {
	greatest := 0
	for _, v := range val {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}
