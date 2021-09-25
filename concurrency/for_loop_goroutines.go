package main

import "fmt"

func main() {
	var i uint64
	inChan := make(chan uint64)
	for i = 1; i < 1000; i++ {
		go func(i uint64) {
			inChan <- i
		}(i)
	}

	for {
		end := false
		select {
		case data, _ := <-inChan:
			fmt.Println(data)
		default:
			fmt.Println("Found exit")
			end = true
			break;
		}

		if end {
			break;
		}
	}
}
