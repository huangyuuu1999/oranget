package main

import (
  "fmt"
  "strconv"
)

type TwoInts struct {
  a int
  b int
}

func main() {
  two1 := new(TwoInts)
  two1.a = 12
  two1.b = 10

  fmt.Printf("two is: %v\n", two1)

  fmt.Println("two is:", two1)
  fmt.Printf("two is: %T\n", two1)
  fmt.Printf("two is: %#v\n", two1)
}


func (tn *TwoInts) String() string {
  return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
