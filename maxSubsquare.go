package main

import (
	"fmt"
)

// Coordinate stores location of a point in a matrix
type Coordinate struct {
	horizontal, vertical int
}

// get minimum value between horizontal and vertical
func min(co Coordinate) int {
	if co.horizontal < co.vertical {
		return co.horizontal
	}

	return co.vertical
}

// calculate the max size of subsquare matrix
// surrounded by 'X'
func subSquareMatrix(input [][]int) int {
	max := 0
	coor := [][]Coordinate{}
	row, col := 0, 0

	for rowIn, rowData := range input {

		coorCol := []Coordinate{}
		for colIn, colData := range rowData {
			tempCoor := Coordinate{horizontal: 0, vertical: 0}

			if colData == 'X' {
				if colIn == 0 {
					tempCoor.horizontal = 1
				} else {
					tempCoor.horizontal = coorCol[colIn-1].horizontal + 1
				}

				if rowIn == 0 {
					tempCoor.vertical = 1
				} else {
					tempCoor.vertical = coor[rowIn-1][colIn].vertical + 1
				}
			}

			coorCol = append(coorCol, tempCoor)
			col = colIn
		}

		coor = append(coor, coorCol)
		row = rowIn
	}

	/*
	   [[{0 0} {0 0} {0 0} {0 0} {1 1}]
	    [{1 1} {0 0} {1 1} {2 1} {3 2}]
	    [{1 2} {0 0} {1 2} {0 0} {1 3}]
	    [{1 3} {2 1} {3 3} {4 1} {5 4}]
	    [{0 0} {0 0} {1 4} {2 2} {3 5}]]
	*/

	// Start from the rightmost bottom most corner
	for i := row; i >= 0; i-- {
		for j := col; j >= 0; j-- {
			// Find smaller of values
			small := min(coor[i][j])

			for small > max {
				if coor[i][(j-small)+1].vertical >= small &&
					coor[(i-small)+1][j].horizontal >= small {
					max = small
				}

				small--
			}
		}
	}

	return max
}

func main() {
	input := [][]int{{'O', 'O', 'O', 'O', 'X'},
		{'X', 'O', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'O', 'X'},
		{'X', 'X', 'X', 'X', 'X'},
		{'O', 'O', 'X', 'X', 'X'},
	}

	fmt.Println(subSquareMatrix(input))
}

// Result: 3
