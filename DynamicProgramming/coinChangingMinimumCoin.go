/*
Given a total and coins of certain denomination with infinite supply, what is the minimum number
of coins it takes to form this total.
*/

package main

import (
	"fmt"
	"math"
)

// Top down approach to get minimum number of coins to get the total.
// It uses a map to store intermediate results.
func minimumCoinTopDown(total int, coins []int, tracker map[int]int) int {

	// if total is 0 then return 0
	if total == 0 {
		return 0
	}

	// if tracker contains minimum value then return same.
	// tracker contains value means it is pre calculated so no need to
	// calculate again
	if val, ok := tracker[total]; ok {
		return val
	}

	// iterate through the coins to check the minimum value
	min := math.MaxInt32
	for i := 0; i < len(coins); i++ {
		// if coin is greater than the total then continue
		if coins[i] > total {
			continue
		}

		val := minimumCoinTopDown(total-coins[i], coins, tracker)

		// if val is lesser than the found min then overwrite min value
		if val < min {
			min = val
		}
	}

	// if min is MaxInt32 then don't change, otherwise add 1 to it
	if min != math.MaxInt32 {
		min += 1
	}

	// memorize the value by add into tracker
	tracker[total] = min

	return min

}

func main() {
	total := 13
	coins := []int{7, 3, 2, 6}
	tracker := make(map[int]int)

	fmt.Println("Top Down result for minimum coin count:", minimumCoinTopDown(total, coins, tracker))
}

/*
Top Down result for minimum coin count: 2
*/
