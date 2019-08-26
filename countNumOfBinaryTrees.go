package main

import (
	"fmt"
)

func countTrees(nodes int) int64 {
	if nodes <= 0 {
		return 0
	}

	results := make([]int64, nodes+1)

	results[0], results[1] = 1, 1

	for i := 2; i <= nodes; i++ {
		for j := 0; j < i; j++ {
			results[i] += results[j] * results[i-j-1]
		}
	}

	return results[nodes]
}

func main() {
	fmt.Println(countTrees(5))
}

// Result : 42
