package main

import (
  "fmt"
)

func main() {
  age := 24
  if age < 18 {
    fmt.Println("The age was less than 18")
  } else if age == 18 {
    fmt.Println("The age was equal to 18.")
  } else {
    fmt.Println("The age was greater than 18")
  }
}
