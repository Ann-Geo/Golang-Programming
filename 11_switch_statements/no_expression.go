package main

import "fmt"

func main() {
	inputstring := "cat"

	switch  {
	case len(inputstring) == 2:
		fmt.Println("Length of the string is 2")
	case inputstring == "dog":
		fmt.Println("Input string is dog")
	case inputstring == "elephant":
		fmt.Println("Input string is elephant")
	default:
		fmt.Println("The is default case")
	}
}
