/*
https://www.geeksforgeeks.org/coin-change-dp-7/

Given a value N, if we want to make change for N cents, and we have infinite supply of
each of S = { S1, S2, .. , Sm} valued coins, how many ways can we make the change? The order of coins does not matter.
For example, for N = 4 and S = {1,2,3}, there are four solutions: {1,1,1,1},{1,1,2},{2,2},{1,3}.
So output should be 4. For N = 10 and S = {2, 5, 3, 6}, there are five solutions:
{2,2,2,2,2}, {2,2,3,3}, {2,2,6}, {2,3,5} and {5,5}. So the output should be 5.
*/

package main

import (
	"fmt"
)

// it returns numbmer of possible solution for
// coin changing problem
func numberOfSolution(total int, coins []int) int {
	T := make([][]int, len(coins)+1)
	for i := 0; i < len(coins)+1; i++ {
		T[i] = make([]int, total+1)
	}

	for i := 0; i < len(T); i++ {
		T[i][0] = 1
	}

	for i := 1; i < len(T); i++ {
		for j := 1; j < len(T[0]); j++ {
			if coins[i-1] > j {
				T[i][j] = T[i-1][j]
			} else {
				T[i][j] = T[i][j-coins[i-1]] + T[i-1][j]
			}
		}
	}

	/*
		[[1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
		 [1 0 0 1 0 0 1 0 0 1 0 0 1 0 0 1]
		 [1 0 0 1 1 0 1 1 1 1 1 1 2 1 1 2]
		 [1 0 0 1 1 0 2 1 1 2 2 1 4 2 2 4]
		 [1 0 0 1 1 0 2 2 1 2 3 2 4 4 4 5]
		 [1 0 0 1 1 0 2 2 1 3 3 2 5 5 4 7]]
	*/

	return T[len(coins)][total]
}

func main() {
	total := 15
	coins := []int{3, 4, 6, 7, 9}
	fmt.Println("numberOfSolution:", numberOfSolution(total, coins))

}
