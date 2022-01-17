package main

import "fmt"

func main() {
	factorialNums := []int{12, 10, 5}
	resultChan := make(chan int)
	for _, num := range factorialNums {
		calcFactorial(num, resultChan)
	}

	for i := 1; i <= len(factorialNums); i++ {
		fmt.Println("Found result: ", <-resultChan)
	}
	close(resultChan)
}

func calcFactorial(num int, ch chan<- int) {
	go func() {
		factorial := 1
		for i := num; i >= 1; i-- {
			factorial *= i
		}
		ch <- factorial
	}()
}