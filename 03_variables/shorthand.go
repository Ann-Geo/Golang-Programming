package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := `Hello World`
	f := 'M'

	fmt.Printf("%v - type: %T\n", a, a)
	fmt.Printf("%v - type: %T\n", b, b)
	fmt.Printf("%v - type: %T\n", c, c)
	fmt.Printf("%v - type: %T\n", d, d)
	fmt.Printf("%v - type: %T\n", e, e)
	fmt.Printf("%v - type: %T\n", f, f)
}
