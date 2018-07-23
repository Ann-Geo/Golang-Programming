package main

import "fmt"

func filter(numbers []int, callback1 func(int) bool) []int {
	var xs []int
	for _, v := range numbers {
		if callback1(v) {
			xs = append(xs, v)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{0, 1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}
