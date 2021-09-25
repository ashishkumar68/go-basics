package main

import (
  "fmt"
  "runtime"
  "sync"
)

// run the program with -race flag to enable the race condition detector.
func main() {
  var wg sync.WaitGroup
  sharedCounter := 0
  goroutines := 100
  fmt.Println("The number of goroutines before forloop", runtime.NumGoroutine())
  var mu sync.Mutex
  wg.Add(goroutines)

  for i := 0; i < goroutines; i++ {
    go func() {
      mu.Lock()
      runtime.Gosched()
      sharedCounter += 1
      mu.Unlock()
      wg.Done()
    }()
    fmt.Println("The number of goroutines now are: ", runtime.NumGoroutine())
  }

  wg.Wait()
  fmt.Println("The number of goroutines now are: ", runtime.NumGoroutine())
  fmt.Println("Counter is finally:", sharedCounter)
}
