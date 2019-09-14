/*
You are given a m x n 2D grid initialized with these three possible values.
 -1 - A wall or an obstacle.
 0 - A gate.
 INF - Infinity means an empty room. We use the value 231 - 1 = 2147483647 to represent INF as
 you may assume that the distance to a gate is less than 2147483647.

 Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, it should be filled with INF
*/

package main

import (
	"container/list"
	"fmt"
	"math"
)

// location stores wall and gate location
type location struct {
	row, col int
}

// it gives direction to go right,left,up and down
// from a location
var dir = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

// stores location information of gates
func gates(w [][]int32) *list.List {
	ll := list.New()
	for i, val := range w {
		for j, val1 := range val {
			if val1 == 0 {
				ll.PushBack(location{row: i, col: j})
			}
		}
	}

	return ll
}

// stores distance of wall from a reachable gate
func wallAndGates(w [][]int32) {
	ll := gates(w)

	for ll.Len() != 0 {
		gt := ll.Front()
		for _, v := range dir {
			nr := gt.Value.(location).row + v[0]
			nc := gt.Value.(location).col + v[1]

			if nr < 0 || nc < 0 || nr >= len(w) || nc >= len(w[0]) || w[nr][nc] != math.MaxInt32 {
				continue
			}

			w[nr][nc] = w[gt.Value.(location).row][gt.Value.(location).col] + 1
			ll.PushBack(location{row: nr, col: nc})
		}
		ll.Remove(gt)
	}

	/*
		[[3 -1 0 1]
		 [2 2 1 -1]
		 [1 -1 2 -1]
		 [0 -1 3 4]]
	*/
	fmt.Println(w)
}

func main() {
	room := [][]int32{
		{math.MaxInt32, -1, 0, math.MaxInt32},
		{math.MaxInt32, math.MaxInt32, math.MaxInt32, -1},
		{math.MaxInt32, -1, math.MaxInt32, -1},
		{0, -1, math.MaxInt32, math.MaxInt32},
	}

	wallAndGates(room)
}

// [[3 -1 0 1] [2 2 1 -1] [1 -1 2 -1] [0 -1 3 4]]
