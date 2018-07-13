package main

import "fmt"

func main() {
	fmt.Println(53)
	fmt.Printf("%d\t %b \n", 53, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d \t %b \t %x \t %q\n", i, i, i, i)
	}
}
