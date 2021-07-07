package main

import (
	"fmt"
)

func main() {
	var count int64
	count = 0
	for count < 1000000000 {
		count += 1
	}
	fmt.Println(count)
}
