package main

import "fmt"

func main(){
  student := make([]string, 1)
  students := make([][]string, 10)
  fmt.Println(student)
	fmt.Println(students)
	fmt.Println(students == nil)
  	student[0] = "Sarah"
	student = append(student, "jenny")
	fmt.Println(student)
}
