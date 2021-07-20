package main

import (
  "fmt"
)

func main() {
  for current := 33; current < 123 ; current++ {
    fmt.Printf("%d is type:%T, str representation: %s",current, current, string(current))
    fmt.Println()
  }
}
