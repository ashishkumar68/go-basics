package main

import (
  "fmt"
)

func main() {
  c := make(chan int)
  go func(c chan int){
    for i := 0; i < 10; i++ {
      c <- i
    }
    close(c)
  }(c)

  for val := range c {
    fmt.Println(val)
  }
  fmt.Println("Done")
}
