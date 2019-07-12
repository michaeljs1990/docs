package main

import (
  "fmt"
)

// Example (4)
// 0 0 0 0
// 0 0 0 1
// 0 1 0 0
// 0 0 0 0
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
    copy(newBoard[i], board[i])
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
