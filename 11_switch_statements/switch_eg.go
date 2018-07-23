package main

import "fmt"

func main() {
	switch 'c' {
	case 'a':
		fmt.Println("The character is 'a'")
	case 'b':
		fmt.Println("The character is 'b'")
	case 'c':
		fmt.Println("The character is 'c'")
	default:
		fmt.Println("The is default case")
	}
}
