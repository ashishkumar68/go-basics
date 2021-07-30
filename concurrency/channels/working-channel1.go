package main

import (
  "fmt"
)

func main()  {
  c := make(chan int)
  go func() {
    c <- 41
  }()
  // this statement blocks until the go routine returns the value.
  fmt.Println(<-c)

}
