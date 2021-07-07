package main

import (
  "fmt"
)

func main() {
  s := `Hello golang playground`;
  byteString := []byte(s)
  fmt.Printf("%T", byteString)
  fmt.Println()
  fmt.Println(byteString)
  fmt.Println()

  for _, val := range s {
    fmt.Printf("Representations for %s", string(val))
    fmt.Println()
    // UTF-8
    fmt.Printf("%#U", val)
    fmt.Println()
    // Hexadecimal
    fmt.Printf("%#X", val)
    fmt.Println()
    // Lowercase Hexadecimal
    fmt.Printf("%#x", val)
    fmt.Println()
  }
}
