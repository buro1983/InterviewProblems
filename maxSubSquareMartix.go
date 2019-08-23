/*
Maximum size square sub-matrix with all 1s
Given a binary matrix, find out the maximum size square sub-matrix with all 1s.
*/

package main

import (
	"fmt"
)

// return minimum vaue
func min(a, b int) int {
	if a < b {
		return a
	} 
	
	return b
}

// pre processing input array to find max 1s
func preProcessing(input [][]int, rows, cols int) int {
	r := rows+1
	c := cols+1
	max := 0
	
	// create additional matrix with rows+1 x cols+1 size
	ppm := make([][]int, r)
	for index := range ppm {
		ppm[index] = make([]int, c)
	}
	
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if input[i-1][j-1] == 1 {
			    // minimum value among the above 
                // side column, diagonally opposite of current position
				ppm[i][j] = min(ppm[i-1][j], min(ppm[i][j-1], ppm[i-1][j-1])) + 1
				
				if ppm[i][j] > max {
					max = ppm[i][j]
				}
			} else {
				ppm[i][j] = 0
			}
		}
	}
	
	/*
	    ppm:
	    [[0 0 0 0 0 0] 
	     [0 0 1 1 0 1] 
	     [0 1 1 0 1 0] 
	     [0 0 1 1 1 0] 
	     [0 1 1 2 2 0] 
	     [0 1 2 2 3 1] 
	     [0 0 0 0 0 0]]
	*/
	
	return max
}

// returns the size of max sub square matrix
func maxSubSquareMatrix(input [][]int) int {

	rows := len(input)
	cols := len(input[0])
	
	return preProcessing(input, rows, cols)
}

func main() {
	input := [][]int{{0, 1, 1, 0, 1},  
                     {1, 1, 0, 1, 0},  
                     {0, 1, 1, 1, 0},  
                     {1, 1, 1, 1, 0},  
                     {1, 1, 1, 1, 1},  
                     {0, 0, 0, 0, 0}}

	fmt.Println("Result:",maxSubSquareMatrix(input))
}

//Result: 3
