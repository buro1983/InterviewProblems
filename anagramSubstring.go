/*
https://www.geeksforgeeks.org/anagram-substring-search-search-permutations/

Given a text txt[0..n-1] and a pattern pat[0..m-1], write a function
search(char pat[], char txt[]) that prints all occurrences of pat[]
and its permutations (or anagrams) in txt[]. You may assume that n > m
Expected time complexity is O(n)
*/

package main

import (
	"fmt"
)

// compare pattern and source string char count
func compare(pat, src map[rune]int) bool {
	if len(pat) != len(src) {
		return false
	}

	for k, v := range pat {
		if v1, ok := src[k]; !ok {
			return false
		} else if v1 != v {
			return false
		}
	}

	return true
}

// print position if anagram present
func anagramSearch(pat, src []rune) {
	patCharCount := make(map[rune]int)
	srcCharCount := make(map[rune]int)

	for _, val := range pat {
		patCharCount[val]++
	}

	for i := 0; i < len(pat); i++ {
		srcCharCount[src[i]]++
	}

	for j := len(pat); j < len(src); j++ {
		if compare(patCharCount, srcCharCount) {
			fmt.Println("Found at index:", j-len(pat))
		}

		srcCharCount[src[j]]++
		srcCharCount[src[j-len(pat)]]--

		// clean source substring map if char count is zero
		if v, ok := srcCharCount[src[j-len(pat)]]; ok {
			if v == 0 {
				delete(srcCharCount, src[j-len(pat)])
			}
		}
	}

	if compare(patCharCount, srcCharCount) {
		fmt.Println("Found at index:", len(src)-len(pat))
	}
}

func main() {
	src := []rune("BACDGABCDA")
	pat := []rune("ABCD")

	anagramSearch(pat, src)
}

/*
Found at index: 0
Found at index: 5
Found at index: 6
*/
