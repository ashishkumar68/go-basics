package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	countUpto := 8000
	var counter int64

	ch := make(chan int64)
	for i := 0; i < countUpto; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			ch <- atomic.LoadInt64(&counter)
		}()
	}
	for i := 0; i < countUpto; i++ {
		<-ch
	}
	fmt.Println("Done")
}
