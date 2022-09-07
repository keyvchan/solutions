package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDistance("sea", "eat"))
	fmt.Println(minDistance("leetcode", "etco"))
	fmt.Println(minDistance("a", "b"))
	fmt.Println(minDistance("park", "spake"))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(word1 string, word2 string, memo map[Key]int) int {
	if word1 == "" || word2 == "" {
		return 0
	}
	if val, ok := memo[Key{word1, word2}]; ok {
		return val
	}

	result := 0
	if word1[0] == word2[0] {
		// we could include this character in the LCS
		result = max(result, lcs(word1[1:], word2[1:], memo)+1)
	} else {
		// we could not include this character in the LCS
		result = max(result, lcs(word1[1:], word2, memo))
		result = max(result, lcs(word1, word2[1:], memo))
	}
	memo[Key{word1, word2}] = result
	return result
}

type Key struct {
	word1 string
	word2 string
}

func minDistance(word1 string, word2 string) int {
	memo := map[Key]int{}
	result := lcs(word1, word2, memo)

	return len(word1) + len(word2) - 2*result
}
