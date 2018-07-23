package main

import "fmt"

func main() {
	array := []float64{1, 2, 3, 4, 5}
	n := average(array...)
	fmt.Println(n)
}

func average(val ...float64) float64 {
	var total float64
	for _, v := range val {
		total += v
	}
	return total / float64(len(val))
}
