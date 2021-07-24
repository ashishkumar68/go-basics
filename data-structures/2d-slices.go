package main

import (
  "fmt"
)

func main() {
  rakesh := []string{"Rakesh", "Singhal", "24", "Pune"}
  sachin := []string{"Sachin", "Paliwal", "28", "Noida"}

  people := [][]string{rakesh, sachin}
  fmt.Println(people)
}
