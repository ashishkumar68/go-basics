package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'findMaximumValue' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY prices
 *  2. INTEGER_ARRAY pos
 *  3. LONG_INTEGER_ARRAY amount
 */
func BuildCumulativePricing(pricesLength int32, prices []int32) ([]int64, int64) {
	var cumulativeSum int64
	cumulativePrices := make([]int64, 0)
	for i := int32(0); i < pricesLength; i++ {
		cumulativeSum += int64(prices[i])
		cumulativePrices = append(cumulativePrices, cumulativeSum)
	}

	return cumulativePrices, cumulativeSum
}

var calculatedPosAmountMaxVal = make(map[string]int32)

func findMaximumValue(prices []int32, pos []int32, amount []int64) []int32 {
	// Write your code here
	ans := make([]int32, 0)
	queries := len(pos)
	pricesLength := int32(len(prices))
	cumulativePricing, sum := BuildCumulativePricing(pricesLength, prices)

	for q := 0; q < queries; q++ {
		availAmount := amount[q]
		index := pos[q] - 1
		hash := fmt.Sprintf("%d_%d", index, availAmount)
		if maxVal, ok := calculatedPosAmountMaxVal[hash]; ok {
			ans = append(ans, maxVal)
		} else {
			maxVal = FindAnswer(index, pricesLength, cumulativePricing, prices, sum, availAmount)
			ans = append(ans, maxVal)
			calculatedPosAmountMaxVal[hash] = maxVal
		}
	}

	return ans
}

func FindAnswer(
	index int32,
	pricesLength int32,
	cumulativePrices []int64,
	prices []int32,
	sum int64,
	availAmount int64,
) int32 {
	currentPrice := int64(prices[index])
	newAvail := availAmount - currentPrice
	if newAvail < 0 {
		return 0
	} else if newAvail == 0 || newAvail < currentPrice || (index+1) == pricesLength || int64(prices[index+1]) > newAvail {
		return 1
	} else {
		if newAvail >= sum {
			return pricesLength - index
		} else if (index + 2) == pricesLength {
			if newAvail <= int64(prices[index+1]) {
				return 2
			} else {
				return 1
			}
		}
	}
	// Binary search for the index until we can purchase.
	currentCumulPrice := cumulativePrices[index]
	beg := index + 1
	end := pricesLength - 1
	mid := (beg + end) / 2
	var foundIndex int32 = -1
	for beg < end {
		effCumulPriceMid := int64(cumulativePrices[mid] - currentCumulPrice)
		if effCumulPriceMid == newAvail {
			foundIndex = mid
			break
		} else if effCumulPriceMid < newAvail {
			nextCumulIndex := mid + 1
			if nextCumulIndex <= end {
				nextEffCumulPriceMid := int64(cumulativePrices[nextCumulIndex] - currentCumulPrice)
				if nextEffCumulPriceMid > newAvail {
					foundIndex = mid
					break
				} else if nextEffCumulPriceMid == newAvail {
					foundIndex = nextCumulIndex
					break
				} else {
					beg = nextCumulIndex
				}
			}
		} else {
			prevCumulIndex := mid - 1
			if prevCumulIndex >= beg {
				prevEffCumulPriceMid := int64(cumulativePrices[prevCumulIndex] - currentCumulPrice)
				if prevEffCumulPriceMid <= newAvail {
					foundIndex = prevCumulIndex
					break
				} else {
					end = prevCumulIndex
				}
			}
		}
		mid = (beg + end) / 2
	}
	if foundIndex == -1 {
		foundIndex = pricesLength - 1
	}

	return foundIndex - index + 1
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	//
	stdout, err := os.Create("/Users/ashishkumar/purchase_optimization_output.txt")
	checkError(err)

	defer stdout.Close()
	file, err := os.Open("/Users/ashishkumar/random_input.txt")
	if err != nil {
		fmt.Println("Could not read from file due to error", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	pricesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var prices []int32

	for i := 0; i < int(pricesCount); i++ {
		pricesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		pricesItem := int32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}

	posCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var pos []int32

	for i := 0; i < int(posCount); i++ {
		posItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		posItem := int32(posItemTemp)
		pos = append(pos, posItem)
	}

	amountCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var amount []int64

	for i := 0; i < int(amountCount); i++ {
		amountItem, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		if err != nil {
			fmt.Println("error on line:", i+1)
		}
		checkError(err)
		amount = append(amount, amountItem)
	}

	result := findMaximumValue(prices, pos, amount)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
