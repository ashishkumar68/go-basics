package main

import (
  "fmt"
)

func main() {
  sum := func(nums ...int) int {
    total := 0
    for _, val := range nums {
      total += val
    }

    return total
  }

  numbersToAdd := []int{34, 24, 43, 22, 45, 84, 727, 21}

  fmt.Println(sum(numbersToAdd...))
}
