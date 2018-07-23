package main

import "fmt"

func main() {
	switch 'c' {
	case 'a', 'b':
		fmt.Println("The character is 'a' or 'b'")
	case 'c', 'd':
		fmt.Println("The character is 'c' or 'd'")
	case 'e':
		fmt.Println("The character is 'e'")
	default:
		fmt.Println("The is default case")
	}
}
