package main

import "fmt"

func main() {
	//transaction := make([]int, 3)
	transactions := make([][]int, 0, 5)

	for i := 0; i < 5; i++ {
		transaction := make([]int, 0, 3)
		for j := 0; j < 3; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	//fmt.Println(transaction)
	fmt.Println(transactions)
}
