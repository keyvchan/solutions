package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
	fmt.Println(longestPalindromeSubseq("a"))
	fmt.Println(longestPalindromeSubseq("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"))
}

func dfs(s string, i int, j int, memo [][]int) int {
	if i > j {
		return 0
	}
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	if i == j {
		memo[i][j] = 1
		return 1
	}
	result := 0
	if s[i] == s[j] {
		result = 2 + dfs(s, i+1, j-1, memo)
	} else {
		result = max(dfs(s, i+1, j, memo), dfs(s, i, j-1, memo))
	}
	memo[i][j] = result

	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindromeSubseq(s string) int {

	memo := [][]int{}
	for i := 0; i < len(s); i++ {
		memo = append(memo, make([]int, len(s)))
	}

	// start from the middle, expand to both sides

	return dfs(s, 0, len(s)-1, memo)
}
