package main

import (
  "fmt"
)

func main() {
  currentVal := 12
  fmt.Println(currentVal)
  fmt.Println(&currentVal)
  changeVal(&currentVal)
  fmt.Println(currentVal)
}

func changeVal(newVal *int) {
  *newVal = 13;
}
