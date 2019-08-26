/*
Maximum sum such that no two elements are adjacent
for array and circular array
*/

package main

import (
	"fmt"
)

// return maximum value
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// returns maximum sum for non adjacent numbers
func maxSum(input []int) int {
	if len(input) == 0 {
		return 0
	}

	incl := input[0]
	excl := 0

	for _, val := range input[1:] {
		temp := incl
		incl = max(excl+val, incl)
		excl = temp
	}

	return incl
}

// returns maximum sum for non adjacent numbers
// in circular array. In circular array 1st elecemt
// is adjacent to last element
func maxSumCircularArray(input []int) int {
	if len(input) == 0 {
		return 0
	}

	if len(input) == 1 {
		return input[0]
	}

	// maxsum excluding 1st element from left to right
	maxSum1 := maxSum(input[1:])

	// maxsum excluding nth element from right to left
	maxSum2 := maxSum(input[:len(input)-1])

	return max(maxSum1, maxSum2)
}

func main() {
	input := []int{5, 5, 10, 100, 10, 5}
	fmt.Println("Result:", maxSum(input))

	input = []int{1, 2, 3, 1}
	fmt.Println("Result CircularArray:", maxSumCircularArray(input))
}

/*
Result: 110
Result CircularArray: 4
*/
