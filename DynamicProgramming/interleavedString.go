/*
Find if a string is interleaved of two other strings
Given three strings A, B and C. Write a function that checks
whether C is an interleaving of A and B. C is said to be
interleaving A and B, if it contains all characters of
A and B and order of all characters in individual
strings are preserved.
*/

package main

import (
	"fmt"
)

// returns true if strC is an interleaving of A and B
// otherwise false
func findInterleavedString(strA, strB, strC string) bool {
	lenStrA := len(strA)
	lenStrB := len(strB)

	if lenStrA+lenStrB != len(strC) {
		return false
	}

	ppm := make([][]bool, lenStrA+1)
	for index := range ppm {
		ppm[index] = make([]bool, lenStrB+1)
	}

	for i := 0; i < len(ppm); i++ {
		for j := 0; j < len(ppm[0]); j++ {
			k := i + j - 1

			if i == 0 && j == 0 {
				ppm[i][j] = true

			} else if i == 0 {
				if string(strC[k]) == string(strB[j-1]) {
					ppm[i][j] = ppm[i][j-1]
				}

			} else if j == 0 {
				if string(strC[k]) == string(strA[i-1]) {
					ppm[i][j] = ppm[i-1][j]
				}

			} else {
				if string(strA[i-1]) == string(strC[k]) {
					ppm[i][j] = ppm[i-1][j]
				} else if string(strB[j-1]) == string(strC[k]) {
					ppm[i][j] = ppm[i][j-1]
				} else {
					ppm[i][j] = false
				}
			}
		}
	}

	/*
	    ppm:
	     |  0    X     X     Z    T
	   --|----------------------------
	   0 |[[true true true false false]
	   X | [true true true true false]
	   X | [true true false true false]
	   Y | [false false false true true]
	   M | [false false false false true]]
	*/

	return ppm[lenStrA][lenStrB]
}

func main() {
	strA := "XXYM"
	strB := "XXZT"
	strC := "XXXZXYTM"

	fmt.Println("Result:", findInterleavedString(strA, strB, strC))
}

// Result: true
