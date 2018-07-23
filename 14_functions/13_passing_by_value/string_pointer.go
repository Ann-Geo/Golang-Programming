package main

import "fmt"

func main() {
	name := "Sarah"
	fmt.Println(&name)
	changeMe(&name)

	fmt.Println(name)
	fmt.Println(&name)
}

func changeMe(z *string) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = "Jane"
	fmt.Println(z)
	fmt.Println(*z)
}
