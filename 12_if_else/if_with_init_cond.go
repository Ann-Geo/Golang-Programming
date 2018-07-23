package main

import "fmt"

func main() {

	b := true

  if food := "Chocolate"; b {
    fmt.Println(food)
  }
  //fmt.Println(food)
  //error: scope of the variable food is inside the if loop only

}
