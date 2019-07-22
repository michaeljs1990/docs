package main

import (
	"fmt"
)

func solver(price []int, size int) int {
	if size == 0 {
		return 0
	}

	var maxCost int
	for i := size; i > 0; i-- {
		cost := price[i-1] + solver(price, size-i)

		if cost > maxCost {
			maxCost = cost
		}
	}

	return maxCost
}

func main() {
	price := []int{1, 5, 8, 9, 10}
	rets := solver(price, 4)
	fmt.Println(rets)
}
