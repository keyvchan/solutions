package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
	fmt.Println(integerBreak(120))
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func dfs(n int, memo map[int]float64) float64 {
	if n == 1 {
		return 1
	}
	if val, ok := memo[n]; ok {
		return val
	}

	result := 0.0
	for i := 1; i < n; i++ {
		result = max(result, dfs(i, memo)*dfs(n-i, memo))
		result = max(result, dfs(i, memo)*float64(n-i))
		result = max(result, float64(i)*dfs(n-i, memo))
		result = max(result, float64(i)*float64(n-i))
	}
	memo[n] = result

	return result

}

func integerBreak(n int) int {
	memo := map[int]float64{}

	return int(math.Mod(dfs(n, memo), 1000000007))

}
