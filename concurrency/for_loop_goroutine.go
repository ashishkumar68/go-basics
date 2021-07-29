package main

import (
  "fmt"
  "runtime"
  "sync"
)

var wg sync.WaitGroup

func loop1() {
  for i := 1; i <= 10; i++ {
    fmt.Println("Loop 1 iteration:", i)
  }
  wg.Done()
}

func loop2() {
  for i := 1; i <= 10; i++ {
    fmt.Println("Loop 2 iteration:", i)
  }
  wg.Done()
}

func main() {
  fmt.Println("sys arch:", runtime.GOARCH)
  fmt.Println("sys OS:", runtime.GOOS)
  fmt.Println("CPUs available:", runtime.NumCPU())
  fmt.Println("Goroutines running:", runtime.NumGoroutine())

  wg.Add(2)
  go loop1()
  go loop2()

  fmt.Println("Goroutines running:", runtime.NumGoroutine())
  wg.Wait()
  fmt.Println("Goroutines running:", runtime.NumGoroutine())
}
