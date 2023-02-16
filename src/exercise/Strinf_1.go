package main

import (
  "fmt"
  "strconv"
)

type T struct {
  a int
  b float32
  c string
}

func main() {
  t := T{2, 3.14, "miemie"}
  fmt.Println(t)
)
}

func (t *T) String() string {
  return "[" + t.c + strconv.Itoa(t.a) + strconv.Itoa(t.b) + "]"
}
 
