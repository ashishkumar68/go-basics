package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func OperationWait(ctx context.Context, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Waiting for ", duration.Seconds(), "Seconds")
	select {
	case <- time.After(duration):
		fmt.Println("Sleep", duration.Seconds(), "sec complete")
	case <- ctx.Done():
		fmt.Println("Context", duration.Seconds(), "sec cancelled")
	}
}

func main() {
	var wg sync.WaitGroup
	newContext, cancel := context.WithCancel(context.TODO())
	wg.Add(4)
	go func() {
		OperationWait(newContext, time.Second * 20, &wg)
		fmt.Println("Cancelling 20 sec context")
		cancel()
		fmt.Println("Cancelled 20 sec context")
	}()
	go func() {
		OperationWait(newContext, time.Second * 120, &wg)
	}()
	go func() {
		OperationWait(newContext, time.Second * 240, &wg)
	}()
	go func() {
		OperationWait(newContext, time.Second * 200, &wg)
	}()

	wg.Wait()
	fmt.Println("About to exit..")
}
