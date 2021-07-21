package main

import (
  "fmt"
)

func main() {
  for init := 0; init < 100; init++ {
    if init % 2 == 0 {
      fmt.Println(init)
    }
  }
}
