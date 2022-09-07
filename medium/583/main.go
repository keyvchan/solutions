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

func lcs(word1 string, word2 string, index1 int, index2 int, memo [][]int) int {
	if index1 == len(word1) || index2 == len(word2) {
		return 0
	}
	if val := memo[index1][index2]; val != 0 {
		return val
	}

	result := 0
	if word1[index1] == word2[index2] {
		// we could include this character in the LCS
		result = max(result, lcs(word1, word2, index1+1, index2+1, memo)+1)
	} else {
		// we could not include this character in the LCS
		result = max(result, lcs(word1, word2, index1+1, index2, memo))
		result = max(result, lcs(word1, word2, index1, index2+1, memo))
	}
	memo[index1][index2] = result
	return result
}

func minDistance(word1 string, word2 string) int {
	memo := [][]int{}
	for i := 0; i < len(word1); i++ {
		memo = append(memo, make([]int, len(word2)))
	}

	result := lcs(word1, word2, 0, 0, memo)

	return len(word1) + len(word2) - 2*result
}
