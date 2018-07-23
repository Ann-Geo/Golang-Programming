package main

import "fmt"

func main() {
	n := average(1, 2, 3, 4, 5)
	fmt.Println(n)
}

func average(val ...float64) float64 {
	fmt.Println(val)
	fmt.Printf("%T\n", val)
	var total float64
	for _, v := range val {
		total += v
	}
	return total / float64(len(val))
}
