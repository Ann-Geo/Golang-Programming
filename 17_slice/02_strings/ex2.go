package main

import "fmt"

func main() {

	places := []string{
		"New York",
		"Mumbai",
		"Tokyo",
		"Bangalore",
		"California",
		"Sydney",
		"Delhi",
	}

	for i, currentEntry := range places {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(places); j++ {
		fmt.Println(places[j])
	}

	fmt.Print("[1:2] ")
	fmt.Println(places[1:2])
	fmt.Print("[:2] ")
	fmt.Println(places[:2])
	fmt.Print("[5:] ")
	fmt.Println(places[5:])
	fmt.Print("[:] ")
	fmt.Println(places[:])
}
