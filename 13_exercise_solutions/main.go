package main

import "fmt"

func main() {
	var name string
	fmt.Println("Hello World")
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)

	var large, small int
	fmt.Println("Enter a large no and a small no: ")
	fmt.Scan(&large, &small)
	fmt.Println("The remainder of large/small = ", large%small)

	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	for j := 1; j <= 100; j++ {
		if j%3 == 0 && j%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if j%3 == 0 {
			fmt.Println("Fizz")
		} else if j%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(j)
		}
	}
	var sum int = 0
	for k := 0; k < 100; k++ {
		if k%3 == 0 || k%5 == 0 {
			sum = sum + k
		}
	}
	fmt.Println(sum)
}
