## Roomba

### Question

This is a path finding problem for a 2d matrix. It likely has some well known name however
I am not familiar with it. If you are please PR and provide a link. Given a 2d matrix you
must find how many spaces you can "clean" on it with a given charge. For it to be a valid
path you must get back to the charger before your battery runs out. A space with a value
of zero means you can clean it where one is a wall or other object that you can't traverse.
When can only move up, down, left, and right. Each move reduces your charge by one. You
can always assume the charger is in the top left corner.

```
[
[0, 0, 0, 1],
[0, 0, 1, 0],
[0. 0. 0. 0],
[1. 0, 0, 0],
]
```

Using the graph above as an example lets look at what the output would look like for a given
charge. For a charge of five your output should be four since you can move through the following
spaces before getting to a charge of zero. The following shows a possible path you can take with
the charge underneath it at each point. Many other valid paths are possible.

```
# (Row, Col, Charge)
(0, 0, 5) -> (1, 0, 4) -> (1, 1, 3) -> (0, 1, 2) -> (0, 0, 1) -> (0, 0, 1)

[
[x, x, 0, 1],
[x, x, 1, 0],
[0. 0. 0. 0],
[1. 0, 0, 0],
]
```

I have marked the points which we covered on the graph with an "x". Additionally you can see
that in the path I followed we waste a turn at the end by sitting in the same place. This is
because it's not possible to go to any more squares without becoming stranded off the charger
with zero charge left.

### Solution 1

This is the first solution that I came up with for this problem and it's super bad but since
I was solving everything else with backtracking here is a backtracking/recursion solution.
This falls apart pretty quick. In my case you can't really do anything with a charge over 13
since this has a time complexity of pow(n, 5).

Additionally since it's go we need to make a copy of the map every time we pass it down so
so we don't need to clean up after outselves in cases where the route is a dead end.
Additionally max is needed for find the max element in a slice.

The cases we have to consider are in the following order

 - The charge is zero and we are at the starting position. Return the distance we traveled.
 - The charge is zero but we aren't at the starting point. This is not valid return zero.
 - Try to stay put. This is for cases where we have left over charge or are stuck at the dock.
 - Try to move up.
 - Try to move down.
 - Try to move left.
 - Try to move right.

Every time we take an action to move we call solver again and do the same thing. Every time
we get to the end of the function we find the path that returned the most distance traveled
and return it.

```
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

func solver(room [][]int, charge, row, col int, visited map[string]bool) int {
        if charge == 0 && row == 0 && col == 0 {
                // We have found a possible winning solution
                return len(visited)
        }

        if charge == 0 {
                return 0
        }

        rets := []int{}

        // Stay put
        key := fmt.Sprintf("%d.%d", row, col)
        if (row) >= 0 && charge != 0 {
                visitcopy := copymap(visited)
                visitcopy[key] = true
                i := solver(room, charge-1, row, col, visitcopy)
                rets = append(rets, i)
        }

        // Move up
        key = fmt.Sprintf("%d.%d", row-1, col)
        if (row-1) >= 0 && charge != 0 && room[row-1][col] != 1 {
                visitcopy := copymap(visited)
                visitcopy[key] = true
                i := solver(room, charge-1, row-1, col, visitcopy)
                rets = append(rets, i)
        }

        // Move down
        key = fmt.Sprintf("%d.%d", row+1, col)
        if (row+1) < len(room) && charge != 0 && room[row+1][col] != 1 {
                visitcopy := copymap(visited)
                visitcopy[key] = true
                i := solver(room, charge-1, row+1, col, visitcopy)
                rets = append(rets, i)
        }

        // Move left
        key = fmt.Sprintf("%d.%d", row, col-1)
        if (col-1) >= 0 && charge != 0 && room[row][col-1] != 1 {
                visitcopy := copymap(visited)
                visitcopy[key] = true
                i := solver(room, charge-1, row, col-1, visitcopy)
                rets = append(rets, i)
        }

        // Move right
        key = fmt.Sprintf("%d.%d", row, col+1)
        if (col+1) < len(room[row]) && charge != 0 && room[row][col+1] != 1 {
                visitcopy := copymap(visited)
                visitcopy[key] = true
                i := solver(room, charge-1, row, col+1, visitcopy)
                rets = append(rets, i)
        }

        return max(rets)
}

func main() {
        room := [][]int{
                []int{0, 0, 1, 0, 0},
                []int{0, 0, 0, 0, 0},
                []int{1, 1, 0, 1, 0},
                []int{0, 0, 1, 0, 0},
                []int{1, 0, 0, 0, 0},
        }

        visited := map[string]bool{}

        o := solver(room, 13, 0, 0, visited)
        fmt.Println(o)
}
```