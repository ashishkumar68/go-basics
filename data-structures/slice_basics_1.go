package main

import "fmt"

func PrintContents(arr []int) {
	fmt.Println("Slice is: ", arr, "with Len:", len(arr), "and Cap:", cap(arr))
}

func main() {
	var intArrDirect = [3]int{1, 2, 3}
	fmt.Println("Array is: ", intArrDirect, "Len is:", len(intArrDirect))

	var intArrDec [3]int
	fmt.Println("Array is: ", intArrDec, "Len is:", len(intArrDec))

	var intArrSliceInit = []int{12, 23, 22}
	PrintContents(intArrSliceInit)

	var intArrSliceMake = make([]int, 10, 20)
	PrintContents(intArrSliceMake)

	intArrSliceMake = intArrSliceMake[:5]
	PrintContents(intArrSliceMake)
	intArrSliceMake = intArrSliceMake[:7]
	PrintContents(intArrSliceMake)

}
