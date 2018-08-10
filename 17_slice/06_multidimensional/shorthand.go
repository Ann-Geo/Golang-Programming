package main

import "fmt"

func main() {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(students == nil)
	//student[0] = "Sarah" - This will not work, have to do append
	student = append(student, "Sarah")
	fmt.Println(student)
}
