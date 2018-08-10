package main

import "fmt"

func main() {

	places := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	places[0] = "New york"
	places[1] = "California"
	places[2] = "Texas"
	//greeting[3] = "Savannah" -- This will not work. Cannot add elements to slice
	//directly beyond length.. Have to append to the slice
	places = append(places, "Savannah")

	//append beyond capacity
	places = append(places, "Dallas")
	places = append(places, "orlando", "Miami")

	fmt.Println(places[6])
	fmt.Println(len(places))
	fmt.Println(cap(places))
}
