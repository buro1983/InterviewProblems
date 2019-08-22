package main

import (
	"fmt"
)

// To preprcess input[M][N]. 
// It creates another 2D array such a way that ppdata[i][j] stores sum 
// of elements from (0,0) to (i,j) 
func preProcess(input [][]int) [][]int {

	row := len(input)
	col := len(input[0])
	 
	//create additional 0th row and 0th column
	ppdata := make([][]int, row+1)
	for index := range ppdata {
		ppdata[index] = make([]int, col+1)
	}
	
	// calculate the sum of second row
	for colIn, rowData := range input[0] {
		ppdata[1][colIn+1] = ppdata[1][colIn] + rowData
	}
	
	// calculate the sum of second column
	for rowIn, colData := range input {
		ppdata[rowIn+1][1] = ppdata[rowIn][1] + colData[0]
	}
	
	//calculate the sum of the remaining column
	for i := 2; i <= row; i++ {
		for j := 2; j <= col; j++ {
			ppdata[i][j] = ppdata[i][j-1] + ppdata[i-1][j] + input[i-1][j-1] - ppdata[i-1][j-1]
		}
	} 
	
	return ppdata
}

// compute sum of submatrix between (tr1, tr2) and (br1, br2) 
// using preprocessed data
func sum(ppdata [][]int, tr1, tr2, br1, br2 int) *int {

    if tr1 < 0 || tr2 < 0 || br1 < 0 || br2 < 0 {
        return nil
    }
    
    // increment locations by 1 as preprocessed matrics size is
    // one great than input matrics
	tr1 += 1
	tr2 += 1
	br1 += 1
	br2 += 1
	
    result := ppdata[br1][br2] - ppdata[tr1-1][br2] - ppdata[br1][tr2-1] + ppdata[tr1-1][tr2-1]
	return &result
}

func main() {

	input := [][]int{{2, 0, -3, 4},
			             {6, 3, 2, -1},
			             {5, 4, 7, 3},
			             {2, -6, 8, 1},}

	/*
        ppdata:
	    [[0 0 0 0 0] 
	     [0 2 2 -1 3] 
	     [0 8 11 10 13] 
	     [0 13 20 26 32] 
	     [0 15 16 30 37]]
	*/
	ppdata := preProcess(input)
	
    // tr1 and tr2 denotes location from top of the matrics
	tr1 := 1
	tr2 := 1
    // br1 and br2 denotes location from bottom of the matrics
	br1 := 3
	br2 := 2
    
    result := sum(ppdata, tr1, tr2, br1, br2)
    if result == nil {
        fmt.Println("Matrics location may not be valid")
    } else {
	    fmt.Println(*result)
    }
}

// Result: 18
