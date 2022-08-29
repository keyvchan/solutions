package main

import "fmt"

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}
	if val, ok := memo[n]; ok {
		return val
	}

	result := 0
	for i := 1; i < n; i++ {
		result = max(result, dfs(i, memo)*dfs(n-i, memo))
		result = max(result, dfs(i, memo)*(n-i))
		result = max(result, i*dfs(n-i, memo))
		result = max(result, i*(n-i))
	}
	memo[n] = result

	return result

}

func integerBreak(n int) int {
	memo := map[int]int{}

	return dfs(n, memo)

}
