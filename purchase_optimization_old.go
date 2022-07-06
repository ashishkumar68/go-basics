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

func findMaximumValue1(prices []int32, pos []int32, amount []int64) []int32 {
	// Write your code here
	ans := make([]int32, 0)
	queries := len(pos)
	pricesLength := int32(len(prices))
	for q := 0; q < queries; q++ {
		availAmount := amount[q]
		maxVal := FindAnswer1(pos[q]-1, pricesLength, prices, availAmount)
		ans = append(ans, maxVal)
	}

	return ans
}

func FindAnswer1(pos, pricesLength int32, prices []int32, availAmount int64) int32 {
	var maxVal int32
	for currentPos := pos; currentPos < pricesLength; currentPos++ {
		currentPrice := int64(prices[currentPos])
		if availAmount >= currentPrice {
			maxVal += 1
			availAmount = availAmount - currentPrice
		} else {
			break
		}
	}

	return maxVal
}

func main() {

	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	//
	stdout, err := os.Create("/Users/ashishkumar/purchase_optimization_output_old.txt")
	checkError1(err)

	defer stdout.Close()
	file, err := os.Open("/Users/ashishkumar/random_input.txt")
	if err != nil {
		fmt.Println("Could not read from file due to error", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	pricesCount, err := strconv.ParseInt(strings.TrimSpace(readLine1(reader)), 10, 64)
	checkError1(err)

	var prices []int32

	for i := 0; i < int(pricesCount); i++ {
		pricesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine1(reader)), 10, 64)
		checkError1(err)
		pricesItem := int32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}

	posCount, err := strconv.ParseInt(strings.TrimSpace(readLine1(reader)), 10, 64)
	checkError1(err)

	var pos []int32

	for i := 0; i < int(posCount); i++ {
		posItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine1(reader)), 10, 64)
		checkError1(err)
		posItem := int32(posItemTemp)
		pos = append(pos, posItem)
	}

	amountCount, err := strconv.ParseInt(strings.TrimSpace(readLine1(reader)), 10, 64)
	checkError1(err)

	var amount []int64

	for i := 0; i < int(amountCount); i++ {
		amountItem, err := strconv.ParseInt(strings.TrimSpace(readLine1(reader)), 10, 64)
		if err != nil {
			fmt.Println("error on line:", i+1)
		}
		checkError1(err)
		amount = append(amount, amountItem)
	}

	result := findMaximumValue1(prices, pos, amount)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine1(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError1(err error) {
	if err != nil {
		panic(err)
	}
}
