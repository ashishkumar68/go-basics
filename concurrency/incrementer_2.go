package main

import "fmt"

func Adder(countCh <-chan int) <-chan int {
	sumChan := make(chan int)
	go func() {
		var sum int
		for val := range countCh {
			sum += val
		}
		sumChan <- sum
		close(sumChan)
	}()

	return sumChan
}

func Incrementer(publishCh chan<- int) <-chan bool {
	completeChan := make(chan bool)
	go func() {
		for i := 0; i < 20; i++ {
			publishCh <- 1
		}

		close(completeChan)
	}()

	return completeChan
}

func main() {
	pc := make(chan int)
	sumChan := Adder(pc)

	cc1 := Incrementer(pc)
	cc2 := Incrementer(pc)

	<-cc1
	<-cc2
	close(pc)
	fmt.Println("Received Sum:", <-sumChan)
}
