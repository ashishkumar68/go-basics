package main

import (
  "fmt"
)

func main() {
  fmt.Println(sum(2, 3, 4, 92, 34, 32))

  func() {
    fmt.Println("this is inside anonymous function")
  }()
  
}

func sum(numbers ...int) int {
  sum := 0;
  for _, val := range numbers {
    sum += val
  }

  return sum
}
