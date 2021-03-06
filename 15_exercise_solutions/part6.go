package main

import "fmt"

func main() {
	var xs []int
	k := 0
	i := 1
	xs = append(xs, i)
	j := 2
	xs = append(xs, j)
	sum := j
	for j < 4000000 {
		k = i + j
		if k > 4000000 {
			break
		}
		if k%2 == 0 {
			sum = sum + k
		}
		xs = append(xs, k)
		i = j
		j = k
	}
	fmt.Println(xs)
	fmt.Println(sum)
}

/*

Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
