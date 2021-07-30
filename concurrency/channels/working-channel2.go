package main

import (
  "fmt"
)

func main() {
  c := make(chan int, 1) // room for 1 value without go-routine
  c <- 32
  // this should work because we have buffer for communicating 1 value.
  fmt.Println(<-c)
}
