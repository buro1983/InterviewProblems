/*
https://www.geeksforgeeks.org/z-algorithm-linear-time-pattern-searching-algorithm/

This algorithm finds all occurrences of a pattern in a text in linear time.
Let length of text be n and of pattern be m, then total time taken is O(m + n) with linear space complexity.
*/

package main

import (
	"fmt"
)

// get Z algo array
func getZ(m []rune, z []int, l int) {

	// left and right window
	left, right, k := 0, 0, 0

	for i := 1; i < l; i++ {
		// if i > right means nothing matches, so calculate
		// z[i] using naive way
		if i > right {
			left, right = i, i
			for right < l && m[right-left] == m[right] {
				right++
			}

			z[i] = right - left
			right--

		} else {
			// k = i-left, k is the number which matches in left,right interval
			k = i - left

			// if z[k] less than remaining interval then z[i] will be equal to z[k]
			if z[k] < right-i+1 {
				z[i] = z[k]
			} else {
				left = i
				for right < l && m[right-left] == m[right] {
					right++
				}

				z[i] = right - left
				right--
			}
		}
	}
}

// run zalgorithm and create z array
func zalgorithm(src, pat string) {
	if len(pat) > len(src) {
		return
	}

	merge := pat + "$" + src
	l := len(merge)

	z := make([]int, l)

	getZ([]rune(merge), z, l)

	for index := range z {
		if z[index] == len(pat) {
			fmt.Println("Pattern found at index:", index-len(pat)-1)
		}
	}
}

func main() {
	src := "aaabcxyzaaaabczaaczabbaaaaaabc"
	pat := "aaabc"

	zalgorithm(src, pat)
}

/*
Pattern found at index: 0
Pattern found at index: 9
Pattern found at index: 25
*/
