package main

import (
	"fmt"
)

func main() {
	var records [][]string
	// student 1
	student1 := make([]string, 4)
	student1[0] = "sarah"
	student1[1] = "James"
	student1[2] = "75.00"
	student1[3] = "50.00"
	// store the record
	records = append(records, student1)
	// student 2
	student2 := make([]string, 4)
	student2[0] = "Jess"
	student2[1] = "Stephen"
	student2[2] = "90.00"
	student2[3] = "78.00"
	// store the record
	records = append(records, student2)
	// print
	fmt.Println(records)
}
