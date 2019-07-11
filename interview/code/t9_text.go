package main

import (
	"fmt"
)

// T9 mapping of number to chars
var t9 = map[int]string{
	1: "",
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
	0: "",
}

var rets = []string{}

func appender(word string) {
	for _, v := range rets {
		if v == word {
			return
		}
	}

	rets = append(rets, word)
}

func solver(number []int, word string, num int) {
	if len(number) == num {
		appender(word)
		return
	}

	if number[num] == 0 || number[num] == 1 {
		solver(number, "", num+1)
		appender(word)
	}

	for _, v := range t9[number[num]] {
		solver(number, word+string(v), num+1)
	}
}

func main() {
	number := []int{8, 2, 4, 2, 2, 0, 5, 5, 5, 5}
	solver(number, "", 0)
	fmt.Println(rets)
}
