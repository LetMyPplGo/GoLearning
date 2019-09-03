package main

import (
  "fmt"
)

func int_seq() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main4() {
  next_int := int_seq()

  fmt.Println(next_int())
  fmt.Println(next_int())

  next_int2 := int_seq()

  fmt.Println(next_int2())
  fmt.Println(next_int())

}
