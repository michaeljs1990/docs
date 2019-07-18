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

// Example of how the solution builds up it's answer for the inputs
// ABC and AC
//
//   0 A C
//  ------
// 0|0 0 0
//
// A|0 1 1
//
// B|0 1 1
//
// C|0 1 2
//
// When we find a match we increment by one and then for computing
// non matches we take the larger of item above it or to the left of it.
func solver(a, b string) int {

  // Create table to hold all calls the +1 is needed for holding
  // the empty string
  memo := make([][]int, len(a)+1)
  for i := range memo {
    memo[i] = make([]int, len(b)+1)
  }

  // Fill in the first column so everything is 0. This may seem a bit strange
  // but it's setting up the case of having an empty string so the rest is easy
  // to setup later. The next loop does the same but for the first row.
  for i := 0; i < len(a); i++ {
    memo[i][0] = 0
  }

  for i := 0; i < len(b); i++ {
    memo[0][i] = 0
  }

  for i := 1; i <= len(a); i++ {
    for j := 1; j <= len(b); j++ {
      if a[i-1] == b[j-1] {
        memo[i][j] = memo[i-1][j-1] + 1
      } else {
        memo[i][j] = Max(memo[i-1][j], memo[i][j-1])
      }
    }
  }

  return memo[len(a)][len(b)]
}

func main() {
	fmt.Println(solver("ABC", "AC"))
}
