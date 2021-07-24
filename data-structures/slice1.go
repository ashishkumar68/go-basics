package main

import (
  "fmt"
)

func main() {
  slice := []int{1, 3, 372, 243, 242}
  appendSlice := []int{4, 6}
  slice = append(slice, appendSlice...)

  fmt.Println(slice)
  // skipping first 2 elements
  fmt.Println(slice[2:])
  // remove 1, and 2th index elements
  slice = append(slice[:1], slice[3:]...)
  fmt.Println(slice)
}
