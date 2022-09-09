package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	currentVal := int64(0)
	numRoutines := 100
	wg.Add(numRoutines)
	for i := 1; i <= numRoutines; i++ {
		go func(num int) {
			atomic.AddInt64(&currentVal, 1)
			fmt.Println(fmt.Sprintf("Goroutine: %d says:", num), atomic.LoadInt64(&currentVal))
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Final counter value:", currentVal)
	fmt.Println("Go os:", runtime.GOOS)
	fmt.Println("Go architecture:", runtime.GOARCH)
}
