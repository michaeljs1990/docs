## N-Queens Problem

### Question

The N queens puzzle is the problem of placing N chess queens on an N x N chessboard 
so that no two queens threaten each other; thus, a solution requires that no two queens
share the same row, column, or diagonal. The N queens puzzle is an example of the
more general n queens problem of placing n non-attacking queens on an n×n chessboard,
for which solutions exist for all natural numbers n with the exception of n = 2 and n = 3.

For example the following board would be a valid configurations because none of the queens
are able to attack the other.

```
0 1 0 0
0 0 0 1
1 0 0 0
0 0 1 0
```

### Solution

You have a few ways to solve this problem with the most common one I have seen being the
use of [backtracking](https://medium.com/@andreaiacono/backtracking-explained-7450d6ef9e1a).
You can also randomly generate a board and then keep moving pieces to the spot with the least
amount of collisions or no collisions at all. Lastly you can come up with a set of heuristics
that will generatea a valid board. The heuristic based version will fall apart though if you
are asked to generate all valid boards and not just one however.

I used recursion and backtracking in the following solution. Additionally I used golang which
isn't a great choice for this problem because of the need to copy the slice every time. For
recursion lets look at what our function definition needs.

```golang 
func solver(size int, board [][]int, cur int) ([][]int, int)
```

For my solution I wanted to keep track of the size of the board, the current state of the board,
and the index of the queen we are trying to place. From this we return the board and the current
index of the queen that was placed. This last bit will make more sense in a second.

Using this we can create a base case that says when the current number of the queen being placed
is equal to size we return the board. Since the cur is an index it being equal to size would mean
all queens have been placed on the board in a valid config and we are ready to exit. If we didn't
meet this case then we loop over the board and try to place a queen at each spot on it and verify
that it is valid. When we find a valid spot for a queen we call the solver function again.... and
again... until we are done.
 
 ```golang
 package main

import (
  "fmt"
)

func makeBoard(size int) [][]int {
  board := make([][]int, size)
  for i := 0; i < size; i++ {
    board[i] = make([]int, size)
  }
  return board
}

func copyBoard(board [][]int) [][]int {
  newBoard := makeBoard(len(board))
  for i := 0; i < len(board); i++ {
    for ii, v := range board[i] {
      newBoard[i][ii] = v
    }
  }
  return newBoard
}

func valid(board [][]int, r, c int) bool {
        for checkRow := 0; checkRow < len(board); checkRow++ {
                if board[r][checkRow] == 1 {
                        return false
                }
        }

        for checkCol := 0; checkCol < len(board); checkCol++ {
                if board[checkCol][c] == 1 {
                        return false
                }
        }

        for rt, ct := r, c; rt < len(board) && ct < len(board); rt, ct = rt+1, ct+1 {
                if board[rt][ct] == 1 {
                        return false
                }
        }

        for rt, ct := r, c; rt < len(board) && ct >= 0; rt, ct = rt+1, ct-1 {
                if board[rt][ct] == 1 {
                        return false
                }
        }

        for rt, ct := r, c; rt >= 0 && ct >= 0; rt, ct = rt-1, ct-1 {
                if board[rt][ct] == 1 {
                        return false
                }
        }

        for rt, ct := r, c; rt >= 0 && ct < len(board); rt, ct = rt-1, ct+1 {
                if board[rt][ct] == 1 {
                        return false
                }
        }

        return true
}

func solver(size int, board [][]int, cur int) ([][]int, int) {
  if size == cur {
    return board, size
  }

  for r, _ := range board {
    for c, _ := range board[r] {
      if valid(board, r, c) {
        bc := copyBoard(board)

        bc[r][c] = 1

        bcf, check := solver(size, bc, cur+1)
        if check == size {
          return bcf, check
        }
      }
    }
  }

  return board, 0
}

func main() {
  board := makeBoard(4)

  out, _ := solver(len(board), board, 0)
  for _, v := range out {
    fmt.Println(v)
  }
}
```

You can likely write the valid function much better than I did if you think of this a bit
more.

### Why this is a bad question

Believe it or not some people are not familiar with chess at all however the rules are easy
enough to explain for one piece. Many people have seen this problem so you likely won't get
good signal from it.

### Why this is a good question

If someone gets through the question fast you can ask them to change it to return all valid
boards. Additionally lots of little optimizations can be made to speed up your solver.