// https://www.techiedelight.com/partition-problem
package main

import (
	"fmt"
)

func SumSlice(x []int) int {
	xSum := 0

	for _, v := range x {
		xSum += v
	}

	return xSum
}

func solver(bucket, x []int, num int) bool {
	if SumSlice(x) == (num / 2) {
		return true
	}

	for idx := range bucket {
    tmpBucket := []int{}
    tmpBucket = append(tmpBucket, bucket[:idx]...)
    tmpBucket = append(tmpBucket, bucket[idx+1:]...)

		tmpX := []int{}
		tmpX = append(x, bucket[idx])
		if solver(tmpBucket, tmpX, num) {
		  return true
		}
	}

	return false
}

func main() {
	bucket := []int{7, 3, 1, 5, 4, 8}
	fmt.Println(solver(bucket, []int{}, SumSlice(bucket)))
}
