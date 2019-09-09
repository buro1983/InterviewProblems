/*
https://www.geeksforgeeks.org/snake-ladder-problem-2/

Given a snake and ladder board, find the minimum number of dice
throws required to reach the destination or last cell from source
or 1st cell. Basically, the player has total control over outcome
of dice throw and wants to find out minimum number of throws
required to reach last cell.
*/

package main

import (
	"container/list"
	"fmt"
)

// queueElem holds position and distance
type queueElem struct {
	pos, dist int
}

// Returns minimum dice throw value to reach destination
func findMinDiceThrowToReachDist(m []int) int {
	visited := make([]bool, len(m))

	l := list.New()
	visited[0] = true
	l.PushBack(queueElem{0, 0})

	final := queueElem{}
	for l.Len() != 0 {
		e := l.Front()
		final = e.Value.(queueElem)
		if final.pos == len(m)-1 {
			break
		}

		l.Remove(e)
		for i := final.pos + 1; i <= final.pos+6 && i < len(m); i++ {
			if visited[i] != true {
				elem := queueElem{}
				visited[i] = true
				elem.dist = final.dist + 1

				if m[i] != -1 {
					elem.pos = m[i]
				} else {
					elem.pos = i
				}

				l.PushBack(elem)
			}
		}
	}

	return final.dist
}

func main() {
	m := make([]int, 30)
	for index := range m {
		m[index] = -1
	}

	// Ladders
	m[2] = 21
	m[4] = 7
	m[10] = 25
	m[19] = 28

	// Snakes
	m[26] = 0
	m[20] = 8
	m[16] = 3
	m[18] = 6

	fmt.Println("Min Dice throws required is:", findMinDiceThrowToReachDist(m))
}

//Min Dice throws required is: 3
