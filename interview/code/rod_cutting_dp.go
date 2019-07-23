package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func solver(price []int, size int) int {
	arr := make([]int, size+1)

	for i := 1; i <= size; i++ {
		for j := 1; j <= i; j++ {
      // 0  0
      // 0  1
      // 1  0
			arr[i] = Max(arr[i], price[j-1]+arr[i-j])
		}
	}

	return arr[size]
}

func main() {
	price := []int{1, 5, 8}
	rets := solver(price, 2)
	fmt.Println(rets)
}
