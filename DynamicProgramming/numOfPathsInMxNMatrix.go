/*
Count all possible paths from top left to the bottom right of a MxN matrix
The problem is to count all the possible paths from top left to bottom
right of a MxN matrix with the constraints that from each cell you can
either move only to right or down
*/

package main

import (
	"fmt"
)

// returns count of possible paths to reach bottom
// most cell from the left topmost cell
func numberOfPaths(m, n int) int {
	ppm := make([][]int, m)
	for index := range ppm {
		ppm[index] = make([]int, n)
	}

	for colIndex := range ppm[0] {
		ppm[0][colIndex] = 1
	}

	for rowIndex := range ppm {
		ppm[rowIndex][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			ppm[i][j] = ppm[i-1][j] + ppm[i][j-1]
		}
	}

	/*
	   ppm:
	   [[1 1 1 1]
	    [1 2 3 4]
	    [1 3 6 10]
	    [1 4 10 20]]
	*/

	return ppm[m-1][n-1]
}

func main() {
	fmt.Println("Result:", numberOfPaths(4, 4))
}

// Result: 20
