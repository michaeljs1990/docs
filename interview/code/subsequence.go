// The problem differs from problem of finding common substrings. Unlike substrings,
// subsequences are not required to occupy consecutive positions within the original sequences.
//
// For example, consider the two following sequences X and Y
//
// X: ABCBDAB
// Y: BDCABA
//
// The length of the LCS is 4
// LCS are BDAB, BCAB, and BCBA

package main

import (
	"fmt"
)

// Yes... it's true golang doesn't have a simple
// max function for ints...
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func solver(a, b string) int {
  if len(a) == 0 || len(b) == 0 {
    return 0
  }

  if a[0] == b[0] {
    return solver(a[1:], b[1:]) + 1
  }

  return Max(solver(a[1:], b), solver(a, b[1:]))
}

func main() {
	fmt.Println(solver("ABBCCC", "ZZC"))
}
