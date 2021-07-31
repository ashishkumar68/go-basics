package main

import (
  "fmt"
)

func send(eve, odd, quit chan<- int) {
  for i := 0; i < 100; i++ {
    if i % 2 == 0 {
      eve <- i
    } else {
      odd <- i
    }
  }
  quit <- 0
}

func receive(eve, odd, quit <-chan int) {
  for {
    select {
      case val := <-eve:
        fmt.Println("eve:", val)
      case val := <-odd:
        fmt.Println("odd:", val)
      case val := <-quit:
        fmt.Println("quit:", val)
        return
    }
  }
}

func main() {
  eve  := make(chan int)
  odd  := make(chan int)
  quit := make(chan int)

  go send(eve, odd, quit)

  receive(eve, odd, quit)

  fmt.Println("About to exit")
}
