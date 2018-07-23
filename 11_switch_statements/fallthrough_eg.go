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
		fallthrough
	case 'd':
		fmt.Println("The character is 'd'")
		fallthrough
	case 'e':
		fmt.Println("The character is 'e'")
	default:
		fmt.Println("The is default case")
	}
}
