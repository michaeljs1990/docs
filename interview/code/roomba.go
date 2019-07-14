package main

import (
	"fmt"
)

func max(rets []int) int {
	biggest := 0
	for _, v := range rets {
		if v > biggest {
			biggest = v
		}
	}

	return biggest
}

func copymap(visited map[string]bool) map[string]bool {
	visitcopy := map[string]bool{}
	for k, v := range visited {
		visitcopy[k] = v
	}

  return visitcopy
}

func solver(room [][]int, charge, row, col int, visited map[string]bool) (int, bool) {
	if charge == 0 && row == 0 && col == 0 {
		// We have found a possible winning solution
		return len(visited), true
	}

	if charge == 0 {
		return 0, false
	}

	rets := []int{}

	// Stay put
	key := fmt.Sprintf("%d.%d", row, col)
	if (row) >= 0 && charge != 0 {
    visitcopy := copymap(visited)
		visitcopy[key] = true
		i, b := solver(room, charge-1, row, col, visitcopy)
		rets = append(rets, i)
		if !b {
			delete(visitcopy, key)
		}
	}

	// Move up
	key = fmt.Sprintf("%d.%d", row-1, col)
	if (row-1) >= 0 && charge != 0  && room[row-1][col] != 1{
    visitcopy := copymap(visited)
		visitcopy[key] = true
		i, b := solver(room, charge-1, row-1, col, visitcopy)
		rets = append(rets, i)
		if !b {
			delete(visitcopy, key)
		}
	}

	// Move down
	key = fmt.Sprintf("%d.%d", row+1, col)
	if (row+1) < len(room) && charge != 0 && room[row+1][col] != 1 {
    visitcopy := copymap(visited)
		visitcopy[key] = true
		i, b := solver(room, charge-1, row+1, col, visitcopy)
		rets = append(rets, i)
		if !b {
			delete(visitcopy, key)
		}
	}

	// Move left
	key = fmt.Sprintf("%d.%d", row, col-1)
	if (col-1) >= 0 && charge != 0 && room[row][col-1] != 1 {
    visitcopy := copymap(visited)
		visitcopy[key] = true
		i, b := solver(room, charge-1, row, col-1, visitcopy)
		rets = append(rets, i)
		if !b {
			delete(visitcopy, key)
		}
	}

	// Move right
	key = fmt.Sprintf("%d.%d", row, col+1)
	if (col+1) < len(room[row]) && charge != 0 && room[row][col+1] != 1{
    visitcopy := copymap(visited)
		visitcopy[key] = true
		i, b := solver(room, charge-1, row, col+1, visitcopy)
		rets = append(rets, i)
		if !b {
			delete(visitcopy, key)
		}
	}

	return max(rets), true
}

func main() {
	room := [][]int{
		[]int{0, 0, 1, 0},
		[]int{0, 0, 0, 0},
		[]int{1, 1, 1, 1},
		[]int{0, 0, 1, 0},
	}

	visited := map[string]bool{}

	o, _ := solver(room, 9, 0, 0, visited)
	fmt.Println(o)
}
