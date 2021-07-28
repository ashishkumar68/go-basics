package main

import (
  "fmt"
)

func main() {
  defer runInEnd()
  fmt.Println("Print first")
  fmt.Println("Print second")
}

func runInEnd() {
  fmt.Println("Print last")
}
