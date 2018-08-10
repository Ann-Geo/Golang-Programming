package main

import "fmt"

func function(x int) (int, bool) {
	var ret1 int
	ret1 = x / 2
	var ret2 bool
	rem := x % 2
	if rem == 0 {
		ret2 = true
	} else {
		ret2 = false
	}
	return ret1, ret2
}
func main() {
	a, b := function(5)
	fmt.Println(a)
	fmt.Println(b)
}
