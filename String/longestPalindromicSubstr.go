/*
https://www.geeksforgeeks.org/longest-palindrome-substring-set-1/

Given a string, find the longest substring which is palindrome. 
For example, if the given string is “forgeeksskeegfor”, the output should be “geeksskeeg”.
*/

package main

import (
    "fmt"
)

// calculate longest palindrome substeing and length
func longestPalindromeSubStrDynamicProgramming(src []rune) (int,[]rune) {
    max := 0
    start := 0
    T := make([][]bool, len(src))

    for index := range T {
        T[index] = make([]bool, len(src))
    }

    for i := 0; i < len(src); i++ {
        T[i][i] = true
    }

    for j := 0; j < len(src)-1; j++ {
        if src[j] == src[j+1] {
            T[j][j+1] = true
            max = 2
            start = j
        }
    }

    for k := 3; k <= len(src); k++ {
        for i := 0; i < len(src)-k+1; i++ {
            j := i+k-1
            if src[i] == src[j] && T[i+1][j-1] {
                T[i][j] = true
                
                if max < k {
                    max = k
                    start = i
                }
            }
        }
    }

    return max, src[start:start+max]
}

func main() {
    src := []rune("forgeeksskeegfor")
    max, subStr := longestPalindromeSubStrDynamicProgramming(src)
    fmt.Println("Result:",max, string(subStr))
}

// Result: 10 geeksskeeg
