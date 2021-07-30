package main

import (
  "fmt"
)

func main() {
  c := make(chan int, 1)
  // this is alright
  c <- 21
  // but not this one because we only have capacity for 1
  c <- 55
  fmt.Println(<-c)
}
