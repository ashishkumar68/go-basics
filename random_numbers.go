package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	pricesCount := 500000
	positionsCount := 200000
	amountsCount := 200000

	file, err := os.Create("/Users/ashishkumar/random_input.txt")
	if err != nil {
		fmt.Println("Could not create file due to error", err)
		os.Exit(1)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	file.WriteString(fmt.Sprintf("%d\n", pricesCount))

	prices := make([]int32, 0)
	for i := 0; i < pricesCount; i++ {
		newPrice := r.Int31n(1000001)
		if i > 0 {
			previousPrice := prices[i-1]
			if previousPrice > newPrice {
				newPrice += previousPrice - newPrice + 1
			}
		}
		prices = append(prices, newPrice)
		file.WriteString(fmt.Sprintf("%d\n", newPrice))
	}

	file.WriteString(fmt.Sprintf("%d\n", positionsCount))
	positions := make([]int32, 0)
	for i := 1; i <= positionsCount; i++ {
		newPos := r.Int31n(500001)
		if newPos == 0 {
			i -= 1
			continue
		}
		positions = append(positions, newPos)
		file.WriteString(fmt.Sprintf("%d\n", newPos))
	}

	file.WriteString(fmt.Sprintf("%d\n", amountsCount))
	amounts := make([]int64, 0)
	for i := 1; i <= amountsCount; i++ {
		newAmt := r.Int63n(1000000000001)
		amounts = append(amounts, newAmt)
		if i != amountsCount {
			file.WriteString(fmt.Sprintf("%d\n", newAmt))
		} else {
			file.WriteString(fmt.Sprintf("%d", newAmt))
		}
	}

	fmt.Println("Done writing to file.")
}
