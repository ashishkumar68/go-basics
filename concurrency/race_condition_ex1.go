package main

import (
  "fmt"
  "runtime"
  "sync"
  "time"
)

var wg sync.WaitGroup
// run the program with -race flag to enable the race condition detector.
func main() {
  sharedCounter := 0
  goroutines := 100
  fmt.Println("The number of goroutines before forloop", runtime.NumGoroutine())
  wg.Add(goroutines)
  for i := 0; i < goroutines; i++ {
    go func() {
      time.Sleep(time.Second)
      runtime.Gosched()
      sharedCounter += 1
      wg.Done()
    }()
    fmt.Println("The number of goroutines now are: ", runtime.NumGoroutine())
  }

  wg.Wait()
  fmt.Println("The number of goroutines now are: ", runtime.NumGoroutine())
  fmt.Println("Counter is finally:", sharedCounter)
}
