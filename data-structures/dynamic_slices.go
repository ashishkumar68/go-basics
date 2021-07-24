package main

import (
  "fmt"
)

func main() {
  dynamicSlice := make([]int, 10, 10)
  fmt.Println(dynamicSlice)
  dynamicSlice[0] = 12
  dynamicSlice[9] = 22
  fmt.Println(len(dynamicSlice))
  fmt.Println(cap(dynamicSlice))
  fmt.Println(dynamicSlice)
  dynamicSlice = append(dynamicSlice, 123)

  fmt.Println(len(dynamicSlice))
  fmt.Println(cap(dynamicSlice))
  fmt.Println(dynamicSlice)
}
