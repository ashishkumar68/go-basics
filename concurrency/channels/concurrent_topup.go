package main

import (
	"fmt"
	"sync"
)

type TopUp struct {
	Id       int
	Currency string
	Value    float32
}

var topUps = []TopUp{
	{Id: 1, Currency: "INR", Value: 100.00},
	{Id: 2, Currency: "INR", Value: 100.00},
	{Id: 3, Currency: "INR", Value: 100.00},
	{Id: 4, Currency: "INR", Value: 100.00},
	{Id: 5, Currency: "INR", Value: 100.00},
	{Id: 6, Currency: "INR", Value: 100.00},
	{Id: 7, Currency: "INR", Value: 100.00},
	{Id: 8, Currency: "INR", Value: 100.00},
	{Id: 9, Currency: "INR", Value: 100.00},
	{Id: 10, Currency: "INR", Value: 100.00},
}

type NewTopUpRequest struct {
	Value           float32
	VoucherReceiver chan<- TopUp
}

var topUpReqChan = make(chan NewTopUpRequest)

func PublishNewTopUpRequest(request NewTopUpRequest) {
	go func() { topUpReqChan <- request }()
}

func ProcessTopUpRequests() {
	for topUpReq := range topUpReqChan {
		topUpIndex := -1
		for idx, topUp := range topUps {
			if topUp.Value == topUpReq.Value {
				topUpIndex = idx
				break
			}
		}
		if topUpIndex == -1 {
			close(topUpReq.VoucherReceiver)
			continue
		}
		newTopUp := topUps[topUpIndex]
		topUps = append(topUps[:topUpIndex], topUps[topUpIndex+1:]...)

		topUpReq.VoucherReceiver <- newTopUp
	}
}

func main() {
	go ProcessTopUpRequests()

	concurrentRequests := 100
	var wg sync.WaitGroup

	wg.Add(concurrentRequests)
	for i := 1; i <= concurrentRequests; i++ {
		go func(routineNum int) {
			defer wg.Done()
			voucherChan := make(chan TopUp)
			PublishNewTopUpRequest(NewTopUpRequest{
				Value:           100.00,
				VoucherReceiver: voucherChan,
			})

			newVoucher, ok := <-voucherChan
			if !ok {
				fmt.Println(fmt.Sprintf("Goroutine %d could not find any voucher.", routineNum))
				return
			}

			fmt.Println(fmt.Sprintf("Goroutine %d found new voucher: %+v", routineNum, newVoucher))
		}(i)
	}

	wg.Wait()

	fmt.Println("Done.")
}
