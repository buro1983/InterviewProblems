/*
https://www.geeksforgeeks.org/kmp-algorithm-for-pattern-searching/

KMP Algorithm for Pattern Searching
Given a text txt[0..n-1] and a pattern pat[0..m-1], write a function
search(char pat[], char txt[]) that prints all occurrences of pat[] in txt[].
You may assume that n > m.
*/

package main

import (
	"fmt"
)

// Compute temporary array to maintain size of suffix which is same as prefix
func createLongestPrefixSuffix(pat []rune) []int {
	lps := make([]int, len(pat))
	index := 0
	for i := 1; i < len(pat); {
		if pat[i] == pat[index] {
			lps[i] = index + 1
			index++
			i++
		} else {
			if index != 0 {
				index = lps[index-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

// Find matched pattern index using KMP algorithm
func kmpAlgorithm(src, pat string) {
	lps := createLongestPrefixSuffix([]rune(pat))

	i := 0
	j := 0
	for i < len(src) {
		if src[i] == pat[j] {
			i++
			j++
		}

		if j == len(pat) {
			fmt.Println("Pattern Found at index:", i-j)
			j = lps[j-1]
		} else if src[i] != pat[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
}

func main() {
	src := "ABABDABACDABABCABABABABDABACDABABCABAB"
	pat := "ABABCABAB"

	kmpAlgorithm(src, pat)

}

/*
Pattern Found at index: 10
Pattern Found at index: 29
*/
