package main

import (
  "fmt"
)

func main() {
  slice := []int{1, 3, 372, 243, 242}
  appendSlice := []int{4, 6}
  slice = append(slice, appendSlice...)

  fmt.Println(slice)
}
