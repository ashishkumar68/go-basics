package main

import (
  "fmt"
)

func main()  {
  c := make(chan int)
  // this says deadlock because no goroutines
  c <- 41
  fmt.Println(<-c)
}
