package main

import (
	"fmt"
)

var name = "Frontend Masters"


func calculateTax(price float32) (float32, float32, string) {
  return price*0.09, price*0.02, ""

}

func calculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
  stateTax = price * 0.09
  cityTax = price * 0.02
  return 
}

func birthday(pointerAge *int) {
  if (*pointerAge>140) {
    panic("Too old to be true")
  }
  fmt.Printf("The point is %v and the value is %v\n", pointerAge, *pointerAge)
  *pointerAge++
}

func main() {
 // cityTax, _, _:= calculateTax(100)
  // fmt.Println(cityTax)

  defer fmt.Println("Bye!")

  age := 22
  birthday(&age)
  fmt.Println(age)


	// function-scoped variables
	PrintData()
}
